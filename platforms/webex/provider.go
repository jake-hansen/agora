package webex

import "github.com/google/wire"

// Provide returns a new WebexActions.
func Provide() *WebexActions {
	return NewWebex()
}

var (
	// ProviderProductionSet provides a WebexActions for use in production.
	ProviderProductionSet = wire.NewSet(Provide)
)
