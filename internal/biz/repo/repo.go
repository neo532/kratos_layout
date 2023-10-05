package repo

import (
	"context"

	"github.com/neo532/kratos_layout/internal/biz/entity"
)

type TransactionDefaultRepo interface {
	Transaction(c context.Context, fn func(c context.Context) (err error)) (err error)
}

// ========== Demo ==========
type DemoRepo interface {
	Create(c context.Context, demo *entity.Demo) (insID int64, err error)
	Get(c context.Context) (rst []*entity.Demo, err error)
}
