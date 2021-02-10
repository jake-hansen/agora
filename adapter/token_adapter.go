package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

// TokenDTOToDomain converts the given Token from DTO representation to domain representation.
func TokenDTOToDomain(token *dto.Token) *domain.Token {
	convertedToken := &domain.Token{Value: token.Value}
	return convertedToken
}

// TokenDomainToDTO converts the given Token from domain representation to DTO representation.
func TokenDomainToDTO(token *domain.Token) *dto.Token {
	convertedToken := &dto.Token{Value: token.Value}
	return convertedToken
}
