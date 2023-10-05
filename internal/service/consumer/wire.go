package consumer

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewDemoConsumer,
)
