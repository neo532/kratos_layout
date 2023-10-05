package consumer

import (
	"context"

	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/neo532/gokit/log"

	"github.com/neo532/kratos_layout/internal/biz"
)

type DemoConsumer struct {
	dm  *biz.DemoUsecase
	log *log.Helper
}

func NewDemoConsumer(
	dm *biz.DemoUsecase,
	logger klog.Logger) *DemoConsumer {
	return &DemoConsumer{
		dm:  dm,
		log: log.NewHelper(logger),
	}
}

func (s *DemoConsumer) Create(c context.Context, message []byte) (err error) {
	return
}
