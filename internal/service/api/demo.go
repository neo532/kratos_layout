package api

import (
	"context"

	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/neo532/gokit/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/neo532/kratos_layout/internal/biz"
	"github.com/neo532/kratos_layout/internal/biz/entity"
	pb "github.com/neo532/kratos_layout/proto/api/demo/v1"
)

type DemoApi struct {
	pb.UnimplementedDemoServer
	uc  *biz.DemoUsecase
	log *log.Helper
	tag string
}

func NewDemoApi(
	uc *biz.DemoUsecase,
	logger klog.Logger,
) *DemoApi {
	return &DemoApi{
		uc:  uc,
		log: log.NewHelper(logger),
		tag: "api.DemoApi",
	}
}

func (a *DemoApi) Create(c context.Context, req *pb.CreateRequest) (reply *emptypb.Empty, err error) {
	dm := &entity.Demo{
		ID:   req.Id,
		Name: req.Name,
	}
	err = a.uc.Create(c, dm)
	return
}
func (a *DemoApi) Get(c context.Context, req *emptypb.Empty) (reply *pb.GetReply, err error) {
	reply = &pb.GetReply{}
	var rs []*entity.Demo
	rs, err = a.uc.GetList(c)
	if err == nil && len(rs) > 0 {
	}
	return
}
