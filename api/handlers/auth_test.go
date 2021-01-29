package handlers_test

import (
	"bytes"
	"encoding/json"
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

func TestAuthHandler_Login(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var mockCredentials = domain.Auth{
			Credentials: &domain.User{
				Username:  "test",
				Password:  "test",
			},
		}
		var mockToken = domain.Auth{
			Token: "test-token",
		}

		mockSimpleAuthService := new(mocks.SimpleAuthService)
		mockSimpleAuthService.On("Authenticate").Return(mockToken, nil)

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		handlers.NewAuthHandler(router.Group("test"), mockSimpleAuthService)

		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(mockCredentials)
		req, err := http.NewRequest("POST", "/test/auth", payloadBuf)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var retrievedToken domain.Auth
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

		mockSimpleAuthService := new(mocks.SimpleAuthService)

		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		handlers.NewAuthHandler(router.Group("test"), mockSimpleAuthService)

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
}
