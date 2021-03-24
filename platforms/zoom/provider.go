package zoom

import "github.com/google/wire"

func Provide() *ZoomActions {
	return NewZoom()
}

var (
	ProviderProductionSet = wire.NewSet(Provide)
)
