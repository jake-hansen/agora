package cookieservicemock

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type CookieService struct {
	mock.Mock
}

func (c *CookieService) SetCookie(g *gin.Context, name string, value string, maxAge int, path string, httpOnly bool) {
	_ = c.Called(g, name, value, maxAge, path, httpOnly)
}
