package usecase

import (
	"context"
	"fmt"
	"net/http"
	"plato-tech/muly/domain"
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

func (uu *userUsecase) Insert(c *gin.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	receivedUser, err := uu.userRepo.Insert(ctx, user)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		return nil
	}
	c.JSON(http.StatusOK, receivedUser)
	return nil
}

func (uu *userUsecase) GetByEmail(c *gin.Context, email string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	receivedUser, err := uu.userRepo.GetByEmail(ctx, email)
	if err != nil {
		c.String(http.StatusNotFound, fmt.Sprintf("error: %s", err))
		return nil
	}
	c.JSON(http.StatusOK, receivedUser)
	return nil
}
func (uu *userUsecase) Update(c *gin.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	receivedUser, err := uu.userRepo.Update(ctx, user)
	if err != nil {
		c.String(http.StatusNotFound, fmt.Sprintf("error: %s", err))
		return nil
	}
	c.JSON(http.StatusOK, receivedUser)
	return nil
}
