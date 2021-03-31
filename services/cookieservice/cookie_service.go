package cookieservice

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CookieService struct {
	Cfg Config
}

type Config struct {
	Domain	string
	SecureCookies bool
	SameSite http.SameSite
}

func NewCookieService(cfg Config) *CookieService {
	return &CookieService{
		Cfg: cfg,
	}
}

func (c *CookieService) SetCookie(g *gin.Context, name string, value string, maxAge int, path string, httpOnly bool) {
	g.SetSameSite(c.Cfg.SameSite)
	g.SetCookie(name, value, maxAge, path, c.Cfg.Domain, c.Cfg.SecureCookies, httpOnly)
}
