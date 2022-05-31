package main

import (
	"context"
	pb "demo/api/demo/v1"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var orderClient pb.DemoClient
var conn *grpc.ClientConn

func main() {
	Init()
	Update("2",0,"true","")
	Close()
}

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("grpc link err:" + err.Error())
	}
	orderClient = pb.NewDemoClient(conn)
}

func Create(orderNo, userName string, amount float32, status, fileUrl string) {
	rsp, err := orderClient.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		Order: &pb.Order {
			OrderNo: orderNo,
			UserName: userName,
			Amount: amount,
			Status: status,
			FileUrl: fileUrl,
		},
	})
	if err != nil {
		panic("grpc 创建订单失败:" + err.Error())
	}
	fmt.Println("Create:" + rsp.Result)
}

func List() {
	rsp, err := orderClient.ListOrder(context.Background(), &pb.ListOrderRequest{})
	if err != nil {
		panic("grpc 查询列表失败:" + err.Error())
	}
	fmt.Println("List:")
	for _, o := range rsp.Orders {
		fmt.Println(o)
	}
}

func Update(orderNo string, amount float32, status, fileUrl string) {
	rsp, err := orderClient.UpdateOrder(context.Background(), &pb.UpdateOrderRequest{
		OrderNo: orderNo,
		Amount: amount,
		Status: status,
		FileUrl: fileUrl,
	})
	if err != nil {
		panic("grpc 更新订单失败:" + err.Error())
	}
	fmt.Println("Update:" + rsp.Result)
}

func Get(orderNo string) {
	rsp, err := orderClient.GetOrder(context.Background(), &pb.GetOrderRequest{OrderNo: orderNo})
	if err != nil {
		panic("grpc 查询失败:" + err.Error())
	}
	fmt.Println(rsp.Order)
}

func Delete(orderNo string) {
	rsp, err := orderClient.DeleteOrder(context.Background(), &pb.DeleteOrderRequest{
		OrderNo: orderNo,
	})
	if err != nil {
		panic("删除订单失败: " + err.Error())
	}
	fmt.Println("Delete:" + rsp.Result)
}

func Close() {
	conn.Close()
}