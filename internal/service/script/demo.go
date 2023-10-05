package script

import (
	"context"

	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/neo532/gokit/log"

	"github.com/neo532/kratos_layout/internal/biz"
)

type DemoScript struct {
	log *log.Helper
	dm  *biz.DemoUsecase
}

func NewDemoScript(
	dm *biz.DemoUsecase,
	logger klog.Logger) *DemoScript {
	return &DemoScript{
		dm:  dm,
		log: log.NewHelper(logger),
	}
}

func (s *DemoScript) Get(c context.Context, args string) (err error) {
	_, err = s.dm.GetList(c)
	return
}
