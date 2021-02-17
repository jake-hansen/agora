package corsmiddleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CORSMiddleware struct {
	Config	cors.Config
}

func (c *CORSMiddleware) HandleCORS() gin.HandlerFunc {
	return cors.New(c.Config)
}


