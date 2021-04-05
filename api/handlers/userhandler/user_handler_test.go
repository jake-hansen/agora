package userhandler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/api/handlers/userhandler"
	"github.com/jake-hansen/agora/api/middleware"
	"github.com/jake-hansen/agora/services/mocks/userservicemock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var DTOMockCredentials = dto.User{
	Firstname: "john",
	Lastname:  "doe",
	Username:  "jdoe",
	Password:  "password",
}

func TestUserHandler_RegisterUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userService := userservicemock.Provide()
		userService.On("Register", mock.Anything).Return(1, nil)

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())

		h := userhandler.Provide(userService)
		_ = h.Register(router.Group("test"))

		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(DTOMockCredentials)
		req, err := http.NewRequest("POST", "/test/users", payloadBuf)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("incorrect-format", func(t *testing.T) {
		expectedBody := "{\"validation errors\":[{\"field\":\"Firstname\",\"reason\":\"required\"}," +
			"{\"field\":\"Lastname\",\"reason\":\"required\"},{\"field\":\"Username\",\"reason\":\"required\"}," +
			"{\"field\":\"Password\",\"reason\":\"required\"}]}"

		userService := userservicemock.Provide()
		userService.On("Register", mock.Anything).Return(1, nil)

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())

		h := userhandler.Provide(userService)
		_ = h.Register(router.Group("test"))

		creds := &dto.Credentials{}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(creds)
		req, err := http.NewRequest("POST", "/test/users", payloadBuf)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedBody, w.Body.String())
	})
}
