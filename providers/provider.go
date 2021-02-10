package providers

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/config"
)

var (
	// ProductionSet contains the dependencies necessary for production.
	ProductionSet = wire.NewSet(config.ProviderSet)
)
