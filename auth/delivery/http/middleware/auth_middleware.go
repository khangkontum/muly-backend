package authHandler

import (
	"net/http"
	authUsecase "plato-tech/muly/auth/usecase"
	"plato-tech/muly/utils/appError"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthHandler struct {
	authUsecase authUsecase.AuthUsecase
}

func NewAuthHandler(au *authUsecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{*au}
}

func (ah *AuthHandler) AuthRequired(c *gin.Context) {
	token, err := jwt.Parse(c.Request.Header.Get("Token"), func (token *jwt.Token) (interface, error ) {
		ah.authUsecase.VerifyJWT(c, token )
	})
	if err != nil {
		appError.AbortWithError(c, http.StatusUnauthorized, appError.ErrUnauthorized.Error())
	}
	else {
		c.Next()
	}
}
