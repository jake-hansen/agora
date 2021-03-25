package authhandler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/api/handlers/authhandler"
	"github.com/jake-hansen/agora/api/middleware"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/mocks/authservicemock"
	"github.com/jake-hansen/agora/services/mocks/cookieservicemock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var domainMockCredentials = domain.Auth{
	Credentials: &domain.Credentials{
		Username: "test",
		Password: "test",
	},
}
var domainMockToken = domain.TokenSet{
	Auth:    domain.Token{
		Value:   "test-token",
	},
	Refresh: domain.Token{
		Value:   "test-token",
	},
}

var DTOMockCredentials = dto.Auth{
	Credentials: &dto.Credentials{
		Username: "test",
		Password: "test",
	},
}

var DTOMockToken = dto.Token{
	Value: "test-token",
}

func TestAuthHandler_Login(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockAuthService := authservicemock.Provide()
		mockAuthService.On("Authenticate", mock.AnythingOfType("domain.Auth")).Return(&domainMockToken, nil)
		mockCookieService := cookieservicemock.Provide()
		mockCookieService.On("SetCookie", mock.AnythingOfType("*gin.Context"), "refresh", "test-token", mock.AnythingOfType("int"), "/", true)

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())

		h := authhandler.Provide(mockAuthService, mockCookieService, nil)
		_ = h.Register(router.Group("test"))

		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(DTOMockCredentials)
		req, err := http.NewRequest("POST", "/test/auth", payloadBuf)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var retrievedToken dto.Token
		json.Unmarshal(w.Body.Bytes(), &retrievedToken)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, DTOMockToken, retrievedToken)
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

		mockAuthService := authservicemock.Provide()

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		h := authhandler.Provide(mockAuthService, nil, nil)
		_ = h.Register(router.Group("test"))

		badRequest := `{}`
		testBadRequest(router, badRequest, "{\"validation errors\":[{\"field\":\"Credentials\","+
			"\"reason\":\"required\"}]}")

		badRequest = `{"credentials": {}}`
		testBadRequest(router, badRequest, "{\"validation errors\":[{\"field\":\"Username\",\"reason\":"+
			"\"required\"},{\"field\":\"Password\",\"reason\":\"required\"}]}")

		badRequest = `{"credentials": {"usernames": "test", "passwords": "test"}}`
		testBadRequest(router, badRequest, "{\"validation errors\":[{\"field\":\"Username\",\"reason\":"+
			"\"required\"},{\"field\":\"Password\",\"reason\":\"required\"}]}")

		badRequest = `{"credentials": {"username": "test", "passwords": "test"}}`
		testBadRequest(router, badRequest, "{\"validation errors\":[{\"field\":\"Password\",\"reason\":"+
			"\"required\"}]}")

		badRequest = `{"credentials": {"usernames": "test", "password": "test"}}`
		testBadRequest(router, badRequest, "{\"validation errors\":[{\"field\":\"Username\",\"reason\":"+
			"\"required\"}]}")

		badRequest = `{"credentials": {"usernames": "test", "password": "test"}`
		testBadRequest(router, badRequest, "{\"error\":\"could not parse request body\"}")
	})

	t.Run("invalid-credentials", func(t *testing.T) {
		mockAuthService := authservicemock.Provide()
		var token *domain.TokenSet = nil
		mockAuthService.On("Authenticate", mock.AnythingOfType("domain.Auth")).Return(token,
			errors.New("username or password not correct"))

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		h := authhandler.Provide(mockAuthService, nil, nil)
		_ = h.Register(router.Group("test"))

		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(domainMockCredentials)
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
		mockAuthService := authservicemock.Provide()
		mockAuthService.On("Deauthenticate", mock.AnythingOfType("domain.Token")).Return(nil)

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		h := authhandler.Provide(mockAuthService, nil, nil)
		_ = h.Register(router.Group("test"))

		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(DTOMockToken)
		req, err := http.NewRequest("DELETE", "/test/auth", payloadBuf)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("failure", func(t *testing.T) {
		mockAuthService := authservicemock.Provide()
		mockAuthService.On("Deauthenticate", mock.AnythingOfType("domain.Token")).Return(errors.New("test error"))

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		h := authhandler.Provide(mockAuthService, nil, nil)
		_ = h.Register(router.Group("test"))

		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(DTOMockToken)
		req, err := http.NewRequest("DELETE", "/test/auth", payloadBuf)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("token-missing", func(t *testing.T) {
		mockAuthService := authservicemock.Provide()

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		h := authhandler.Provide(mockAuthService, nil, nil)
		_ = h.Register(router.Group("test"))

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
