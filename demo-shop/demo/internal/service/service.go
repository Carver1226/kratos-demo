package service

import (
	v1 "demo/api/demo/v1"
	"demo/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewDemoService)

type DemoService struct {
	v1.UnimplementedDemoServer

	oc *biz.OrderUsecase
	log *log.Helper
}

func NewDemoService(oc *biz.OrderUsecase, logger log.Logger) *DemoService {
	return &DemoService{oc: oc, log: log.NewHelper(logger)}
}