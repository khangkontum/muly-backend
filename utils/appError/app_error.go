package appError

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrDuplicateEmail = errors.New("duplicate email")
	ErrEditConflict   = errors.New("edit conflict")
	ErrUnauthorized   = errors.New("unauthorized")
)

func AbortWithError(c *gin.Context, statusCode int, err string) {
	c.AbortWithStatusJSON(statusCode, gin.H{"status": false, "message": err})
}
