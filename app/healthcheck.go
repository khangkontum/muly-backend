package main

import (
	// "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthCheckHandler(c *gin.Context) {
	data := map[string]string{
		"status":      "avaible",
		"environment": app.config.env,
		"version":     version,
	}
	err := app.writeJSON(c.Writer, http.StatusOK, envelope{"data": data}, nil)
	if err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(c.Writer, c.Request, err)
	}
}
