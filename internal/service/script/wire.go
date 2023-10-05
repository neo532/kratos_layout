package script

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewDemoScript,
)
