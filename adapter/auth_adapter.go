package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

func AuthDTOToDomain(auth *dto.Auth) *domain.Auth {
	convertedAuth := &domain.Auth{Credentials: &domain.User{
		Username: auth.Credentials.Username,
		Password: auth.Credentials.Password,
	}}
	return convertedAuth
}

func AuthDomainToDTO(auth *domain.Auth) *dto.Auth  {
	convertedAuth := &dto.Auth{Credentials: &dto.Credentials{
		Username: auth.Credentials.Username,
		Password: auth.Credentials.Password,
	}}
	return convertedAuth
}
