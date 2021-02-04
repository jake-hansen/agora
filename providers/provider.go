package providers

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/config"
)

var (
	ProductionSet = wire.NewSet(config.ProviderSet)
)
