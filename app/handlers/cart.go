package handlers

import (
	"context"
	"github.com/arrim/go-cart/pb"
)

func (s *RouteGuideServer) Add(ctx context.Context, request *pb.AddRequest) (*pb.StatusMessage, error) {
	return &pb.StatusMessage{
		Ok:      true,
		Code:    0,
		Message: "ok",
	}, nil
}

func (s *RouteGuideServer) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.StatusMessage, error) {
	return &pb.StatusMessage{
		Ok:      true,
		Code:    0,
		Message: "ok",
	}, nil
}

func (s *RouteGuideServer) Get(ctx context.Context, request *pb.Cart) (*pb.GetResponse, error) {
	return &pb.GetResponse{
		Cart:     nil,
		Products: nil,
		Quantity: 0,
		Amount:   0,
	}, nil
}
