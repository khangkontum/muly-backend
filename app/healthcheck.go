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
	c.JSON(http.StatusOK, map[string]interface{}{"data": data})
}
