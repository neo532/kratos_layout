package data

import (
	"context"

	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/neo532/gokit/database/orm"
	"github.com/neo532/gokit/log"

	"github.com/neo532/kratos_layout/internal/biz/repo"
)

type TransactionDefaultRepo struct {
	db  *orm.Orms
	log *log.Helper
}

func NewTransactionDefaultRepo(defaultDB DatabaseDefault, logger klog.Logger) repo.TransactionDefaultRepo {
	return &TransactionDefaultRepo{
		db:  defaultDB,
		log: log.NewHelper(logger),
	}
}

func (r *TransactionDefaultRepo) Transaction(c context.Context, fn func(ctx context.Context) error) (err error) {
	err = r.db.Transaction(c, fn)
	return
}
