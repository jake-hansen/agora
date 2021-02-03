package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

func AuthDTOToDomain(auth *dto.Auth) *domain.Auth {
	convertedAuth := &domain.Auth{Credentials: UserDTOToDomain(auth.Credentials)}
	return convertedAuth
}

func AuthDomainToDTO(auth *domain.Auth) *dto.Auth  {
	convertedAuth := &dto.Auth{Credentials: UserDomainToDTO(auth.Credentials)}
	return convertedAuth
}
