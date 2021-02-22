package authmiddleware_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/api/middleware"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/services/mocks/authservicemock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthMiddleware_HandleAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockAuthService := authservicemock.Provide()
		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		endpoint := router.Group("test")
		authMiddleware := authmiddleware.Provide(mockAuthService, authmiddleware.ProvideAuthorizationHeaderParser())
		endpoint.Use(authMiddleware.HandleAuth())

		endpoint.GET("", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		mockAuthService.On("IsAuthenticated", mock.Anything).Return(true, nil)

		req, err := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer test")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("failure-no-token", func(t *testing.T) {
		mockAuthService := authservicemock.Provide()
		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		endpoint := router.Group("test")
		authMiddleware := authmiddleware.Provide(mockAuthService, authmiddleware.ProvideAuthorizationHeaderParser())
		endpoint.Use(authMiddleware.HandleAuth())

		endpoint.GET("", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, err := http.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, "{\"error\":\"token not found\"}", w.Body.String())
	})

	t.Run("failure-invalid-token", func(t *testing.T) {
		mockAuthService := authservicemock.Provide()
		router := gin.Default()
		router.Use(middleware.PublicErrorHandler())
		endpoint := router.Group("test")
		authMiddleware := authmiddleware.Provide(mockAuthService, authmiddleware.ProvideAuthorizationHeaderParser())
		endpoint.Use(authMiddleware.HandleAuth())

		endpoint.GET("", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		mockAuthService.On("IsAuthenticated", mock.Anything).Return(false, errors.New("test"))

		req, err := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer test")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, "{\"error\":\"the provided token is not valid\"}", w.Body.String())
	})
}
