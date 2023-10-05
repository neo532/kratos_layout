package server

import (
	"context"

	"github.com/IBM/sarama"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/neo532/gokit/middleware"
	"github.com/neo532/gokit/middleware/server"
	"github.com/neo532/gokit/queue"
	"github.com/neo532/gokit/queue/kafka/consumergroup"

	"github.com/neo532/kratos_layout/internal/conf"
	"github.com/neo532/kratos_layout/internal/service/consumer"
)

// NewConsumer new a HTTP server.
func NewConsumerDefault(
	c context.Context,
	bs *conf.Bootstrap,
	logging klog.Logger,
	dm *consumer.DemoConsumer,
) (csm queue.Consumer) {
	return queue.NewConsumers(newConsumer(c, bs.General, bs.Data.ConsumerDefault.Conf, logging, dm.Create)...)
}

func newConsumer(
	c context.Context,
	cg *conf.General,
	cfg *conf.Data_Consumer,
	logging klog.Logger,
	fn func(c context.Context, message []byte) (err error)) (cs []queue.Consumer) {

	connect := func(c context.Context,
		d *conf.Data_Consumer,
		csm *conf.Data_ConsumerCsm,
		logging klog.Logger,
		fn func(c context.Context, message []byte) (err error),
	) *consumergroup.ConsumerGroup {
		return consumergroup.NewGroup(
			csm.Name,
			csm.Addrs,
			csm.Group,
			consumergroup.WithLogger(logging, c),
			consumergroup.WithTopics(csm.Topics...),
			consumergroup.WithSlowLog(d.MaxSlowtime.AsDuration()),
			consumergroup.WithAutoCommit(true),
			consumergroup.WithBalanceStrategy(sarama.BalanceStrategySticky),
			consumergroup.WithContext(c),
			consumergroup.WithCallback(func(ctx context.Context, message []byte) (err error) {
				ctx = server.SetEntryForCtx(ctx, middleware.EntryConsumer)
				ctx = server.SetEnvForCtx(ctx, cg.Env)
				err = fn(ctx, message)
				return
			}),
		)
	}

	cs = make([]queue.Consumer, 0, 3)
	cs = append(cs, connect(c, cfg, cfg.Default, logging, fn))
	if cfg.Shadow != nil {
		cs = append(cs, connect(c, cfg, cfg.Shadow, logging, fn))
	}
	if cfg.Gray != nil {
		cs = append(cs, connect(c, cfg, cfg.Gray, logging, fn))
	}
	return
}
