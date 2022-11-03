package authHandler

import (
	"net/http"
	"plato-tech/muly/domain"
	"plato-tech/muly/utils/helpers"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(uu domain.UserUsecase) *UserHandler {
	return &UserHandler{uu}
}

func (h *UserHandler) RegisterUserHandler(c *gin.Context) {
	// Read the input Data from API
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// Check if the form is in valid form
	err := helpers.ReadJSON(c, &input)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	// Parse the JSON object form to app domain.User
	user := &domain.User{
		Name:      input.Name,
		Email:     input.Email,
		Activated: false,
	}
	err = user.Password.Set(input.Password)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	h.userUsecase.Insert(c, user)
}
