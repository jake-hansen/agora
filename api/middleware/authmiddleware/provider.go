package authmiddleware

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

// Provide returns an AuthMiddleware configured with the provided AuthService and ParsTokenFunc.
func Provide(authService domain.AuthService, parseToken ParseTokenFunc) *AuthMiddleware {
	return New(&authService, parseToken)
}

// ProvideAuthorizationHeaderParser returns a ParseTokenFunc which parses the Authorization
// header for a Bearer token.
func ProvideAuthorizationHeaderParser() ParseTokenFunc {
	return getTokenFromBearerHeader
}

var (
	// ProviderNoParserSet AuthMiddleware with no AuthService or ParseTokenFunc
	ProviderNoParserSet = wire.NewSet(Provide)

	// ProviderAuthorizationParserSet AuthMiddleware with authorization header parser and no AuthService
	ProviderAuthorizationParserSet = wire.NewSet(Provide, ProvideAuthorizationHeaderParser)

	// ProviderTestSet AuthMiddleware with no AuthService or ParseTokenFunc
	ProviderTestSet = wire.NewSet(Provide)
)
