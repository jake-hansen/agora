package refreshtokenservicemock

import "github.com/stretchr/testify/mock"

func Provide() *RefreshTokenService {
	return &RefreshTokenService{mock.Mock{}}
}

