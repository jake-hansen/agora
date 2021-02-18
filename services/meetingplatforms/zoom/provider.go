package zoom

import "github.com/google/wire"

func Provide() *Zoom {
	return NewZoom()
}

var (
	ProviderProductionSet = wire.NewSet(Provide)
)
