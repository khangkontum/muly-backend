package main

import (
	authHandler "plato-tech/muly/auth/delivery/http"
	"plato-tech/muly/auth/repository/authPostgres"
	authUsecase "plato-tech/muly/auth/usecase"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	// router := httprouter.New()
	// router := mux.NewRouter()
	router := gin.Default()
	router.GET("/v1/healthcheck", app.healthCheckHandler)

	userRepo := authPostgres.NewUserRepo(app.conn)
	userUseCase := authUsecase.NewUserUsecase(userRepo, app.config.timeout)
	authHandler.NewUserHandler(router, userUseCase)

	return router
}
