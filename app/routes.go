package main

import (
	"github.com/gin-gonic/gin"

	_userRepo "plato-tech/muly/auth/repository/postgres/user_repo"
)

func (app *application) routes() *gin.Engine {
	// router := httprouter.New()
	// router := mux.NewRouter()
	router := gin.Default()
	router.GET("/v1/healthcheck", app.healthCheckHandler)

	userRepo := _userRepo.New

	return router
}
