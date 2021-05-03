package domain

import "github.com/gin-gonic/gin"

// CookieService manages operations on HTTP cookies.
type CookieService interface {
	SetCookie(g *gin.Context, name string, value string, maxAge int, path string, httpOnly bool)
}
