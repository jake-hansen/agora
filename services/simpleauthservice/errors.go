package simpleauthservice

type RefreshTokenRevoked struct{}
type RefreshTokenReuse struct{}

func NewRefreshTokenRevokedError() RefreshTokenRevoked {
	return RefreshTokenRevoked{}
}

func (r RefreshTokenRevoked) Error() string {
	return "the refresh token was revoked"
}

func NewRefreshTokenReuseError() RefreshTokenReuse {
	return RefreshTokenReuse{}
}

func (r RefreshTokenReuse) Error() string {
	return "the refresh token was reused"
}
