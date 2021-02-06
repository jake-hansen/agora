package userhandler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jake-hansen/agora/adapter"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
	"net/http"
)

type UserHandler struct {
	UserService *domain.UserService
}

func (u *UserHandler) Register(parentGroup *gin.RouterGroup) error {
	userGroup := parentGroup.Group("user")
	{
		userGroup.POST("", u.RegisterUser)
	}
	return nil
}

func validateHelper(err error) error {
	var verr validator.ValidationErrors
	if err != nil && !errors.As(err, &verr) {
		err = api.NewAPIError(http.StatusBadRequest, err, "could not parse request body")
	}
	return err
}

func(u *UserHandler) RegisterUser(c *gin.Context)  {
	var user dto.User
	err := c.ShouldBind(&user)
	err = validateHelper(err)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	// Register user
	_, err = (*u.UserService).Register(adapter.UserDTOToDomain(&user))
	if err != nil {
		apiError := api.NewAPIError(http.StatusInternalServerError, err, "error occurred during registration")
		_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
	} else {
		c.Status(http.StatusOK)
	}
}

