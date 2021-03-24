package cookieservice

import "github.com/gin-gonic/gin"

type CookieService struct {
	Cfg Config
}

type Config struct {
	Domain	string
	SecureCookies bool
}

func NewCookieService(cfg Config) *CookieService {
	return &CookieService{
		Cfg: cfg,
	}
}

func (c *CookieService) SetCookie(g *gin.Context, name string, value string, maxAge int, path string, httpOnly bool) {
	g.SetCookie(name, value, maxAge, path, c.Cfg.Domain, c.Cfg.SecureCookies, httpOnly)
}
