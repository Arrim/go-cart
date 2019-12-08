package handlers

import (
	"context"
	"github.com/arrim/go-cart/app/repository"
	"github.com/arrim/go-cart/pb"
)

type RouteGuideServer struct {
	repository *repository.Repository
}

func NewRouterGuideServer(r *repository.Repository) *RouteGuideServer {
	return &RouteGuideServer{
		repository: r,
	}
}

func (s *RouteGuideServer) Health(ctx context.Context, empty *pb.Empty) (*pb.HealthMessage, error) {
	return &pb.HealthMessage{Message: "ok"}, nil
}
