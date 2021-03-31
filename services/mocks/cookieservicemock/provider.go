package cookieservicemock

import (
	"github.com/stretchr/testify/mock"
)

func Provide() *CookieService {
	return &CookieService{Mock: mock.Mock{}}
}
