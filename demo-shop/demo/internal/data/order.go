package data

import (
	"context"
	"demo/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type orderRepo struct {
	data *Data
	log *log.Helper
}

type Order struct {
	gorm.Model
	OrderNo string `gorm:"primary_key"`
	UserName string
	Amount float32
	Status string
	FileUrl string
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo {
		data: data,
		log: log.NewHelper(logger),
	}
}

func (or *orderRepo) CreateOrder(ctx context.Context, o *biz.Order) error {
	order := Order {
		OrderNo: o.OrderNo,
		UserName: o.UserName,
		Amount: o.Amount,
		Status: o.Status,
		FileUrl: o.FileUrl,
	}
	if err := or.data.db.Create(&order).Error; err != nil {
		return err
	}
	return nil
}

func (or *orderRepo) GetOrderByOrderNo(ctx context.Context, orderNo string) (*biz.Order, error) {
	var o Order
	err := or.data.db.Where("order_no=?", orderNo).First(&o).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.NotFound("order", "not found by orderNo")
	}
	if err != nil {
		return nil, err
	}
	return &biz.Order {
		OrderNo: o.OrderNo,
		UserName: o.UserName,
		Amount: o.Amount,
		Status: o.Status,
		FileUrl: o.FileUrl,
	}, nil
}

func (or *orderRepo) UpdateOrder(ctx context.Context, orderNo string, amount float32, status, fileUrl string) error {
	var o Order
	if err := or.data.db.Where("order_no=?", orderNo).First(&o).Error; err != nil {
		return err
	}
	if err := or.data.db.Model(&o).Updates(Order{
			Amount: amount,
			Status: status,
			FileUrl: fileUrl,
		}).Error; err != nil {
		return err
	}
	return nil
}

func (or *orderRepo) DeleteOrder(ctx context.Context, orderNo string) error {
	var o Order
	if err := or.data.db.Where("order_no=?", orderNo).Delete(&o).Error; err != nil {
		return err
	}
	return nil
}

func (or *orderRepo) ListOrder(ctx context.Context) ([]*biz.Order, error) {
	var orders []Order
	if err := or.data.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	bo := make([]*biz.Order, len(orders))
	for i, x := range orders {
		bo[i] =  &biz.Order {
			OrderNo: x.OrderNo,
			UserName: x.UserName,
			Amount: x.Amount,
			Status: x.Status,
			FileUrl: x.FileUrl,
		}
	}
	return bo, nil
}