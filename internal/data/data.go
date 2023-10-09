package data

import (
	"context"

	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/neo532/apitool/transport/http/xhttp/client"
	"github.com/neo532/gofr/tool"
	"github.com/neo532/gokit/database/orm"
	"github.com/neo532/gokit/database/redis"
	lredis "github.com/neo532/gokit/lock/distributed/redis"

	"github.com/neo532/kratos_layout/internal/conf"
	sample "github.com/neo532/kratos_layout/proto/client/sample/v1"
)

type (
	DatabaseDefault *orm.Orms
	RedisLock       *redis.Rediss
	//ProducerDefault *queue.Producers
)

// ========== Database ==========
func NewDatabaseDefault(c context.Context, bs *conf.Bootstrap, logger klog.Logger) (DatabaseDefault, func(), error) {
	dbs := newDatabase(c, bs.Data.DatabaseDefault.Conf, logger)
	return dbs, dbs.Cleanup(), dbs.Err
}

// ========== /Database ==========

// ========== Redis ==========

func NewRedisLock(c context.Context, bs *conf.Bootstrap, logger klog.Logger) (RedisLock, func(), error) {
	rdbs := newRedis(c, bs.General, bs.Data.RedisLock.Conf, logger)
	return rdbs, rdbs.Cleanup(), rdbs.Err
}

func NewToolDistributedLock(rdb RedisLock) *tool.DistributedLock {
	return tool.NewDistributedLock(&lredis.GoRedis{Rdb: rdb})
}

// ========== /Redis ==========

// ========== Producer ==========
// func NewProducerDefault(c context.Context, conf *conf.Data_ProducerDefault, logger klog.Logger) (ProducerDefault, func(), error) {
// 	pdcs := newProducer(c, conf.Conf, logger)
// 	return pdcs, pdcs.CleanUp(), pdcs.Err
// }
// ========== /Producer ==========

// ========== Client ==========
func NewSampleXHttpClient(clt client.Client, bs *conf.Bootstrap) (xclt *sample.SampleXHttpClient) {
	xclt = sample.NewSampleXHttpClient(clt)
	xclt.Domain = bs.Third.Sample.Domain
	xclt.WithMiddleware(sample.Demo())
	return
}

// ========== /Client ==========
