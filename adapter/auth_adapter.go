package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

// AuthDTOToDomain converts the given Auth from DTO representation to domain representation.
func AuthDTOToDomain(auth *dto.Auth) *domain.Auth {
	convertedAuth := &domain.Auth{Credentials: &domain.Credentials{
		Username: auth.Credentials.Username,
		Password: auth.Credentials.Password,
	}}
	return convertedAuth
}
