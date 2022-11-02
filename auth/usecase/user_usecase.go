package authUsecase

import (
	"context"
	"net/http"
	"plato-tech/muly/domain"
	"plato-tech/muly/utils/appError"
	"time"

	"github.com/gin-gonic/gin"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(ur domain.UserRepository, contextTimeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		ur,
		contextTimeout,
	}
}

func (uu *userUsecase) Insert(c *gin.Context, user *domain.User) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	receivedUser, err := uu.userRepo.Insert(ctx, user)
	if err != nil {
		appError.AbortWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, receivedUser)
}

func (uu *userUsecase) GetByEmail(c *gin.Context, email string) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	receivedUser, err := uu.userRepo.GetByEmail(ctx, email)
	if err != nil {
		appError.AbortWithError(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, receivedUser)
}
func (uu *userUsecase) Update(c *gin.Context, user *domain.User) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	receivedUser, err := uu.userRepo.Update(ctx, user)
	if err != nil {
		appError.AbortWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, receivedUser)
}
