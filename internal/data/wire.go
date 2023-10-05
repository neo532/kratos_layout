package data

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewDatabaseDefault,
	NewRedisLock,
	NewToolDistributedLock,

	NewTransactionDefaultRepo,
	NewDemoRepo,
)
