  package data

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewTransactionDefaultRepo,
	NewDatabaseDefault,
	NewRedisLock,
	NewToolDistributedLock,
	NewSampleXHttpClient,
	NewDemoRepo,
)
