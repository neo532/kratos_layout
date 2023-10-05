package data

import (
	"context"

	"github.com/neo532/gokit/database/orm"

	"github.com/neo532/kratos_layout/internal/biz/entity"
	"github.com/neo532/kratos_layout/internal/biz/repo"
	"github.com/neo532/kratos_layout/internal/data/db"
)

type DemoRepo struct {
	db *orm.Orms
}

func NewDemoRepo(
	db DatabaseDefault,
) repo.DemoRepo {
	return &DemoRepo{
		db: db,
	}
}

func (r *DemoRepo) Create(c context.Context, d *entity.Demo) (insID int64, err error) {

	data := &db.Demo{
		ID:   d.ID,
		Name: d.Name,
	}
	err = r.db.Write(c).
		WithContext(c).
		Create(data).
		Error

	insID = data.ID
	return
}

func (r *DemoRepo) Get(c context.Context) (rs []*entity.Demo, err error) {
	rs = make([]*entity.Demo, 0, 5)

	var ds []*db.Demo
	err = r.db.Read(c).
		WithContext(c).
		Select("id", "name").
		Order("id desc").
		Find(&ds).
		Error
	if err != nil {
		return
	}
	for _, v := range ds {
		rs = append(rs, &entity.Demo{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return
}
