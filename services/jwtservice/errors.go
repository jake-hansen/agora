package jwtservice

var (
	ErrInvalidTokenPair *InvalidTokenPair = &InvalidTokenPair{}
)

type InvalidTokenPair struct {}

func (i InvalidTokenPair) Error() string {
	return "invalid token pair"
}

func (i InvalidTokenPair) Is(tgt error) bool {
	_, ok := tgt.(InvalidTokenPair)
	return ok
}

