package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type config struct {
	port    int
	env     string
	timeout time.Duration
	db      struct {
		dsn          string
		maxOpenCons  int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}

func importConfig(cfg *config) error {
	var err error
	cfg.port = viper.GetInt("server.address")
	cfg.timeout, err = time.ParseDuration(viper.GetString("timeout"))
	if err != nil {
		log.Fatal("TIME_OUT wrong format")
		return err
	}

	host := viper.GetString("database.host")
	user := viper.GetString("database.user")
	pass := viper.GetString("database.pass")
	name := viper.GetString("database.name")
	address := viper.GetString("database.address")
	sslmode := viper.GetString("database.sslmode")
	cfg.db.dsn = fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=%s", host, user, pass, address, name, sslmode)

	cfg.env = viper.GetString("environment")
	cfg.db.maxOpenCons = viper.GetInt("database.max_open_conns")
	cfg.db.maxIdleConns = viper.GetInt("database.max_idle_conns")
	cfg.db.maxIdleTime, err = time.ParseDuration(viper.GetString("database.max_idle_time"))
	if err != nil {
		fmt.Print(err)
		log.Fatal("MAX_IDLE_TIME wrong format")
		return err
	}
	return nil
}
