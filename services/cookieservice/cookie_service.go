package cookieservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CookieService manages operations on cookies.
type CookieService struct {
	Cfg Config
}

// Config represents configuration for a CookieService.
type Config struct {
	Domain        string
	SecureCookies bool
	SameSite      http.SameSite
}

// NewCookieService returns a CookieService configured with the provided Config.
func NewCookieService(cfg Config) *CookieService {
	return &CookieService{
		Cfg: cfg,
	}
}

// SetCookie a cookie containing the provided values.
func (c *CookieService) SetCookie(g *gin.Context, name string, value string, maxAge int, path string, httpOnly bool) {
	g.SetSameSite(c.Cfg.SameSite)
	g.SetCookie(name, value, maxAge, path, c.Cfg.Domain, c.Cfg.SecureCookies, httpOnly)
}
