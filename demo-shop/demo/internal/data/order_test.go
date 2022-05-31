package data

import (
	"context"
	"demo/internal/biz"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
)

type OrderRepoTestSuite struct {
	suite.Suite
	or *orderRepo
	data *Data
}

func InitMySql() *gorm.DB{
	dsn := "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local"
	Db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	return Db
}

//测试初始化
func (o *OrderRepoTestSuite)SetupTest() {
	d := &Data{db: InitMySql()}
	or := &orderRepo{
		data: d,
	}
	o.or = or
	d.db.Unscoped().Where("1=1").Delete(&Order{})
}

func TestOrderRepoTestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepoTestSuite))
}

func (o *OrderRepoTestSuite)Test_orderRepo() {
	var order *biz.Order
	var orders []*biz.Order
	err := o.or.CreateOrder(context.Background(), &biz.Order{
		OrderNo: "1",
		UserName: "carver",
		Amount:   100,
		Status:   "true",
	})
	o.NoError(err)
	err = o.or.CreateOrder(context.Background(), &biz.Order{
		OrderNo: "2",
		UserName: "mia",
		Amount:   100,
		Status:   "true",
	})
	o.NoError(err)
	order, err = o.or.GetOrderByOrderNo(context.Background(), "1")
	o.NoError(err)
	tOrder := &biz.Order{
		OrderNo: "1",
		UserName: "carver",
		Amount:   100,
		Status:   "true",
	}
	o.Equal(order, tOrder)
	orders, err = o.or.ListOrder(context.Background())
	o.NoError(err)
	tOrders := []*biz.Order {
		tOrder,
		{
			OrderNo: "2",
			UserName: "mia",
			Amount:   100,
			Status:   "true",
		},
	}
	o.Equal(orders, tOrders)
	err = o.or.UpdateOrder(context.Background(), "1", 0, "false", "")
	o.NoError(err)
	order, err = o.or.GetOrderByOrderNo(context.Background(), "1")
	tOrder.Status = "false"
	o.Equal(order, tOrder)
	err = o.or.DeleteOrder(context.Background(), "2")
	o.NoError(err)
}