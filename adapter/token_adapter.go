package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

// TokenDTOToDomain converts the given AuthToken from DTO representation to domain representation.
func TokenDTOToDomain(token *dto.Token) *domain.AuthToken {
	convertedToken := &domain.AuthToken{Value: domain.TokenValue(token.Value)}
	return convertedToken
}

// TokenDomainToDTO converts the given AuthToken from domain representation to DTO representation.
func TokenDomainToDTO(token *domain.AuthToken) *dto.Token {
	convertedToken := &dto.Token{Value: string(token.Value)}
	return convertedToken
}
