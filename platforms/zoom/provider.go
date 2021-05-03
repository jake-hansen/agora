package zoom

import "github.com/google/wire"

// Provide returns a ZoomActions.
func Provide() *ZoomActions {
	return NewZoom()
}

var (
	// ProviderProductionSet provides a ZoomActions for use in production.
	ProviderProductionSet = wire.NewSet(Provide)
)
