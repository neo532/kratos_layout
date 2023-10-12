package biz

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/neo532/gofr/tool"

	"github.com/neo532/kratos_layout/internal/biz/entity"
	"github.com/neo532/kratos_layout/internal/biz/repo"
)

type DemoUsecase struct {
	tx repo.TransactionDefaultRepo
	lk *tool.DistributedLock
	dm repo.DemoRepo
}

func NewDemoUsecase(
	tx repo.TransactionDefaultRepo,
	lk *tool.DistributedLock,
	dm repo.DemoRepo,
) *DemoUsecase {
	return &DemoUsecase{
		tx: tx,
		lk: lk,
		dm: dm,
	}
}

func (uc *DemoUsecase) Create(c context.Context, dm *entity.Demo) (err error) {

	key := strings.Replace(entity.LOCK_DEMO, "{demoId}", strconv.FormatInt(dm.ID, 10), -1)
	var code string
	if code, err = uc.lk.Lock(c, key, 30*time.Second, 2*time.Second); err != nil {
		return
	}
	defer uc.lk.UnLock(c, key, code)

	err = uc.tx.Transaction(c, func(ctx context.Context) (err error) {

		// get
		if _, err = uc.dm.Get(ctx); err != nil {
			return
		}

		// create
		if _, err = uc.dm.Create(ctx, dm); err != nil {
			return
		}

		return
	})

	return
}

func (uc *DemoUsecase) GetList(c context.Context) (rs []*entity.Demo, err error) {
	return uc.dm.Get(c)
}
