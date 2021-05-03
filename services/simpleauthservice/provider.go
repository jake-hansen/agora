package simpleauthservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/jwtservice"
)

// Provide returns a new SimpleAuthService configured with the provided JWTService, UserService, and RefreshTokenService.
func Provide(jwtService jwtservice.JWTService, userService domain.UserService, refreshTokenService domain.RefreshTokenService) *SimpleAuthService {
	return &SimpleAuthService{
		jwtService:          jwtService,
		userService:         userService,
		refreshTokenService: refreshTokenService,
	}
}

var (
	// ProviderProductionSet provides a new SimpleAuthService for use in production.
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.AuthService), new(*SimpleAuthService)))
)
