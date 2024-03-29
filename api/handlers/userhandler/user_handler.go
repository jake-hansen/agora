package userhandler

import (
	"errors"
	"net/http"

	"github.com/jake-hansen/agora/api/middleware/authmiddleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jake-hansen/agora/adapter"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

// UserHandler is the handler that manages operations on Users for the API.
type UserHandler struct {
	UserService    *domain.UserService
	AuthMiddleware *authmiddleware.AuthMiddleware
}

// Register creates 3 endpoints to manage Users.
// /        (POST) - RegisterUser
// /        (GET)  - GetAllUsers
// /:userid (GET)  - GetUser
func (u *UserHandler) Register(parentGroup *gin.RouterGroup) error {
	userGroup := parentGroup.Group("users")
	authenticatedUserGroup := parentGroup.Group("users")
	authenticatedUserGroup.Use(u.AuthMiddleware.HandleAuth())
	{
		userGroup.POST("", u.RegisterUser)
		userGroup.GET("/:userid", u.GetUser)
		authenticatedUserGroup.GET("", u.GetAllUsers)
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

// RegisterUser attempts to register the given user retrieved from the body as JSON.
func (u *UserHandler) RegisterUser(c *gin.Context) {
	var user dto.User
	err := c.ShouldBind(&user)
	err = validateHelper(err)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	// Register user
	createdUserID, err := (*u.UserService).Register(adapter.UserDTOToDomain(&user))
	if err != nil {
		apiError := api.NewAPIError(http.StatusInternalServerError, err, "error occurred during registration")
		_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
	} else {
		resource := &dto.Resource{ID: int(createdUserID)}
		c.JSON(http.StatusCreated, resource)
	}
}

// GetUser gets a user.
func (u *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("userid")

	if userID != "me" {
		err := errors.New("cannot get info about other users")
		apiError := api.NewAPIError(http.StatusBadRequest, err, "cannot get info about other users")
		_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
		return
	}

	user, err := (*u.AuthMiddleware).GetUser(c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, adapter.UserDomainToDTO(user))
}

// GetAllUsers gets all users.
func (u *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := (*u.UserService).GetAll()
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	var userList []*dto.User
	for _, user := range users {
		userList = append(userList, adapter.UserDomainToDTO(user))
	}

	c.JSON(http.StatusOK, userList)
}
