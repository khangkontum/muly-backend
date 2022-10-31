package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// router := httprouter.New()
	router := mux.NewRouter()

	// router.NotFound = http.HandlerFunc(app.notFoundResponse)
	// router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandleFunc("/v1/healthcheck", app.healthCheckHandler).Methods("GET")

	return router
}
