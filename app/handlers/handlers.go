package handlers

import (
	"context"
	"github.com/arrim/go-cart/pb"
)

type RouteGuideServer struct{}

func (s *RouteGuideServer) Health(ctx context.Context, empty *pb.Empty) (*pb.HealthMessage, error) {
	return &pb.HealthMessage{Message: "ok"}, nil
}
