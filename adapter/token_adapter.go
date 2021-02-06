package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

func TokenDTOToDomain(token *dto.Token) *domain.Token {
	convertedToken := &domain.Token{Value: token.Value}
	return convertedToken
}

func TokenDomainToDTO(token *domain.Token) *dto.Token {
	convertedToken := &dto.Token{Value: token.Value}
	return convertedToken
}
