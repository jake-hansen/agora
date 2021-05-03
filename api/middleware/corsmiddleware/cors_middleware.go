package corsmiddleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware serves as middleware to specify CORS settings.
type CORSMiddleware struct {
	Config cors.Config
}

// HandleCORS provides the HandlerFunc to apply CORS settings to requests.
func (c *CORSMiddleware) HandleCORS() gin.HandlerFunc {
	return cors.New(c.Config)
}
