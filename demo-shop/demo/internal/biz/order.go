package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Order struct {
	OrderNo string
	UserName string
	Amount float32
	Status string
	FileUrl string
}

type OrderRepo interface {
	CreateOrder(ctx context.Context, order *Order) error
	GetOrderByOrderNo(ctx context.Context, orderNo string) (*Order, error)
	UpdateOrder(ctx context.Context, orderNo string, Amount float32, Status, FileUrl string) error
	DeleteOrder(ctx context.Context, orderNo string) error
	ListOrder(ctx context.Context) ([]*Order, error)
}

type OrderUsecase struct {
	or OrderRepo
	log *log.Helper
}

func NewOrderUsecase(or OrderRepo, logger log.Logger) *OrderUsecase{
	return &OrderUsecase {
		or: or,
		log: log.NewHelper(logger),
	}
}

func (ou *OrderUsecase) Create(ctx context.Context, o *Order) error {
	return ou.or.CreateOrder(ctx, o)
}

func (ou *OrderUsecase) GetOrder(ctx context.Context, orderNo string) (*Order, error) {
	return ou.or.GetOrderByOrderNo(ctx, orderNo)
}

func (ou *OrderUsecase) UpdateOrder(ctx context.Context, orderNo string, amount float32, status, fileUrl string) error {
	return ou.or.UpdateOrder(ctx, orderNo, amount, status, fileUrl)
}

func (ou *OrderUsecase) DeleteOrder(ctx context.Context, orderNo string) error {
	return ou.or.DeleteOrder(ctx, orderNo)
}

func (ou *OrderUsecase) ListOrder(ctx context.Context) ([]*Order, error) {
	return ou.or.ListOrder(ctx)
}