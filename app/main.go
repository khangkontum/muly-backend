package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const version = "1.0.0"

type application struct {
	config config
	logger *log.Logger
	conn   *sql.DB
}

func main() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Print(err)
		log.Fatal("Error loading .env file")
	}
	var cfg config
	err = importConfig(&cfg)
	if err != nil {
		log.Fatal("Error loading .env content")
	}
	// Initialize a new logger which writes msg to the std,
	// prefixed with the current time and date.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Call the openDB() helper function to create the connection pool,
	// passing in the config struct. If error, log it and exit the app
	// immediately.
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	// Defer a call to db.Close() so that the connection pool is closed before
	// the main() func exits.
	defer db.Close()

	// Log a msg to say that the conn has been successfully established.
	logger.Printf("database connection pool established")

	// Create AWS bucket session
	if cfg.aws.accessKey != "" && cfg.aws.secretAccressKey != "" {
		sess, err := session.NewSessionWithOptions(session.Options{
			Profile: "default",
			Config: aws.Config{
				Region: aws.String(cfg.aws.region),
			},
		})

		if err != nil {
			fmt.Printf("Failed to initialize new session: %v", err)
			return
		}

		s3Client := s3.New(sess)

		err = creadBucket(s3Client, cfg.aws.publicBucketName)
		if err != nil {
			fmt.Printf("Couldn't create bucket: %v", err)
			return
		}
	}

	// Declare an instance of the application struct, containing the config
	// struct and the logger.
	app := &application{
		config: cfg,
		logger: logger,
		conn:   db,
	}

	// Declare a new servrmux and add a /v1/healthcheck route which dispatches
	// requests to the healthCheckHandler method.

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.port),
		Handler:        app.routes(),
		IdleTimeout:    time.Minute,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func creadBucket(client *s3.S3, bucketName string) error {
	_, err := client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})

	return err
}

func openDB(cfg config) (*sql.DB, error) {
	// Create an empty conn pool, using the DSN from the cfg struct.
	// connection pool is safe for concurrent access.
	// first param is driver name, second is data source name

	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetMaxOpenConns(cfg.db.maxOpenCons)
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	// create a context with 5-second timeout deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use PingContext() to establish new conn to the db,
	// passing in the ctx aboved as a param. If conn couldn't be established within
	// 5 secons then return an error.
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
