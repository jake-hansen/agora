package simpleauthservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/jwtservice"
)

// Provide returns a new SimpleAuthService which uses the given JWTService for generating and validating
// JWTs.
func Provide(jwtService jwtservice.JWTService, userService domain.UserService) *SimpleAuthService {
	return &SimpleAuthService{jwtService: jwtService, userService: userService}
}

var (
	// ProviderProductionSet provides a new SimpleAuthService for use in production.
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.AuthService), new(*SimpleAuthService)))
)
