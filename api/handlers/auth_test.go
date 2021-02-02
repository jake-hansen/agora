package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/api/domain"
	"github.com/jake-hansen/agora/api/handlers"
	"github.com/jake-hansen/agora/api/middleware"
	"github.com/jake-hansen/agora/api/services/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var mockCredentials = domain.Auth{
	Credentials: &domain.User{
		Username:  "test",
		Password:  "test",
	},
}
var mockToken = domain.Token{
	Value: "test-token",
}

func TestAuthHandler_Login(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockAuthService := new(mocks.AuthService)
		mockAuthService.On("Authenticate").Return(&mockToken, nil)

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		handlers.NewAuthHandler(router.Group("test"), mockAuthService)

		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(mockCredentials)
		req, err := http.NewRequest("POST", "/test/auth", payloadBuf)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var retrievedToken domain.Token
		json.Unmarshal(w.Body.Bytes(), &retrievedToken)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, mockToken, retrievedToken)
	})

	t.Run("bad-request", func(t *testing.T) {
		testBadRequest := func(router *gin.Engine, badRequest string, validationError string) {
			req, err := http.NewRequest("POST", "/test/auth", strings.NewReader(badRequest))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Equal(t, validationError, w.Body.String())
		}

		mockAuthService := new(mocks.AuthService)

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		handlers.NewAuthHandler(router.Group("test"), mockAuthService)

		badRequest := `{}`
		testBadRequest(router, badRequest, "{\"validation errors\":[{\"field\":\"Credentials\"," +
			"\"reason\":\"required\"}]}")

		badRequest = `{"credentials": {}}`
		testBadRequest(router, badRequest, "{\"validation errors\":[{\"field\":\"Username\",\"reason\":" +
			"\"required\"},{\"field\":\"Password\",\"reason\":\"required\"}]}")

		badRequest = `{"credentials": {"usernames": "test", "passwords": "test"}}`
		testBadRequest(router, badRequest, "{\"validation errors\":[{\"field\":\"Username\",\"reason\":" +
			"\"required\"},{\"field\":\"Password\",\"reason\":\"required\"}]}")

		badRequest = `{"credentials": {"username": "test", "passwords": "test"}}`
		testBadRequest(router, badRequest, "{\"validation errors\":[{\"field\":\"Password\",\"reason\":" +
			"\"required\"}]}")

		badRequest = `{"credentials": {"usernames": "test", "password": "test"}}`
		testBadRequest(router, badRequest, "{\"validation errors\":[{\"field\":\"Username\",\"reason\":" +
			"\"required\"}]}")

		badRequest = `{"credentials": {"usernames": "test", "password": "test"}`
		testBadRequest(router, badRequest, "{\"error\":\"could not parse request body\"}")
	})

	t.Run("invalid-credentials", func(t *testing.T) {
		mockAuthService := new(mocks.AuthService)
		var token *domain.Token = nil
		mockAuthService.On("Authenticate").Return(token,
			errors.New("username or password not correct"))

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		handlers.NewAuthHandler(router.Group("test"), mockAuthService)

		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(mockCredentials)
		req, err := http.NewRequest("POST", "/test/auth", payloadBuf)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, "{\"error\":\"the provided credentials could not be validated\"}", w.Body.String())
	})
}

func TestAuthHandler_Logout(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockAuthService := new(mocks.AuthService)
		mockAuthService.On("Deauthenticate").Return(nil)

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		handlers.NewAuthHandler(router.Group("test"), mockAuthService)

		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(mockToken)
		req, err := http.NewRequest("DELETE", "/test/auth", payloadBuf)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("failure", func(t *testing.T) {
		mockAuthService := new(mocks.AuthService)
		mockAuthService.On("Deauthenticate").Return(errors.New("test error"))

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		handlers.NewAuthHandler(router.Group("test"), mockAuthService)

		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(mockToken)
		req, err := http.NewRequest("DELETE", "/test/auth", payloadBuf)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("token-missing", func(t *testing.T) {
		mockAuthService := new(mocks.AuthService)

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		handlers.NewAuthHandler(router.Group("test"), mockAuthService)

		badRequest := `{}`
		req, err := http.NewRequest("DELETE", "/test/auth", strings.NewReader(badRequest))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "{\"validation errors\":[{\"field\":\"Value\",\"reason\":\"required\"}]}", w.Body.String())
	})
}
