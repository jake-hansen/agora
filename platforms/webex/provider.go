package webex

import "github.com/google/wire"

func Provide() *WebexActions {
	return NewWebex()
}

var (
	ProviderProductionSet = wire.NewSet(Provide)
)
