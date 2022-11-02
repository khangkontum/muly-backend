package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	// router := httprouter.New()
	// router := mux.NewRouter()
	router := gin.New()
	router.GET("/v1/healthcheck", app.healthCheckHandler)

	// router.NotFound = http.HandlerFunc(app.notFoundResponse)
	// router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	return router
}
