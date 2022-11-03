package authUsecase

import (
	"net/http"
	"plato-tech/muly/domain"
	"plato-tech/muly/utils/appError"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthUsecase struct {
	secretKey string
}

func NewAuthHandler(secretKey string) *AuthUsecase {
	return &AuthUsecase{secretKey: secretKey}
}

func (a AuthUsecase) generateJWT(user *domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().AddDate(0, 0, 10)
	claims["authorized"] = true
	claims["id"] = user.ID
	tokenString, err := token.SignedString(a.secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (a AuthUsecase) VerifyJWT(c *gin.Context, token *jwt.Token) {
	_, ok := token.Method.(*jwt.SigningMethodECDSA)
	if !ok {
		appError.AbortWithError(c, http.StatusUnauthorized, appError.ErrUnauthorized.Error())
		return
	}
	if !token.Valid {
		appError.AbortWithError(c, http.StatusUnauthorized, appError.ErrUnauthorized.Error())
		return
	}
	return
}
