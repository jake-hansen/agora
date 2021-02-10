package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

// UserDTOToDomain converts the given User from DTO representation to domain representation.
func UserDTOToDomain(user *dto.User) *domain.User {
	resultUser := &domain.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Username:  user.Username,
		Password:  domain.NewPassword(user.Password),
	}
	return resultUser
}

// UserDomainToDTO converts the given User from domain representation to DTO representation.
// Note that the returned representation will NEVER contain the password hash.
func UserDomainToDTO(user *domain.User) *dto.User {
	resultUser := &dto.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Username:  user.Username,
		Password:  "",
	}
	return resultUser
}

