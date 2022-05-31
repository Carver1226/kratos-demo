package service

import (
	"context"
	"demo/internal/biz"

	pb "demo/api/demo/v1"
)

func convertOrder(order *biz.Order) *pb.Order {
	return &pb.Order {
		OrderNo: order.OrderNo,
		UserName: order.UserName,
		Amount: order.Amount,
		Status: order.Status,
		FileUrl: order.FileUrl,
	}
}
func (s *DemoService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.ResultReply, error) {
	if err := s.oc.Create(ctx, &biz.Order{
		OrderNo: req.Order.OrderNo,
		UserName: req.Order.UserName,
		Amount: req.Order.Amount,
		Status: req.Order.Status,
		FileUrl: req.Order.FileUrl,
	}); err != nil {
		return &pb.ResultReply {
			Result: err.Error(),
		}, err
	}
	return &pb.ResultReply {
		Result: "success",
	}, nil
}
func (s *DemoService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.ResultReply, error) {
	if err := s.oc.UpdateOrder(ctx, req.OrderNo, req.Amount, req.Status, req.FileUrl); err != nil {
		return &pb.ResultReply {
			Result: err.Error(),
		}, err
	}
	return &pb.ResultReply{
		Result: "success",
	}, nil
}
func (s *DemoService) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.ResultReply, error) {
	if err := s.oc.DeleteOrder(ctx, req.OrderNo); err != nil {
		return &pb.ResultReply {
			Result: err.Error(),
		}, err
	}
	return &pb.ResultReply{
		Result: "success",
	}, nil
}
func (s *DemoService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderReply, error) {
	order, err := s.oc.GetOrder(ctx, req.OrderNo)
	if err != nil {
		return nil, err
	}
	return &pb.OrderReply{
		Order: convertOrder(order),
	}, nil
}
func (s *DemoService) ListOrder(ctx context.Context, req *pb.ListOrderRequest) (*pb.ListOrderReply, error) {
	orders , err := s.oc.ListOrder(ctx)
	if err != nil {
		return nil, err
	}
	po := make([]*pb.Order, len(orders))
	for i, x := range orders {
		po[i] = convertOrder(x)
	}
	return &pb.ListOrderReply{
		Orders: po,
	}, nil
}