package server

import (
	"context"
	"fmt"
	"github.com/arrim/go-cart/app/handlers"
	"github.com/arrim/go-cart/app/repository"
	"github.com/arrim/go-cart/pb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

type Server struct {
	wg      sync.WaitGroup
	server  *handlers.RouteGuideServer
	options Options
}

type Options struct {
	GrpcPort int
	RESTPort int
}

func NewServer(options Options, r *repository.Repository) *Server {
	return &Server{
		server:  handlers.NewRouterGuideServer(r),
		options: options,
	}
}

func (s *Server) Start() {
	var err error
	s.wg.Add(1)

	go func() {
		err = s.startGRPC()
		s.wg.Done()

		if err != nil {
			log.Println("gRpc server not run", err)
		}
	}()

	s.wg.Add(1)

	go func() {
		err = s.startREST()
		s.wg.Done()

		if err != nil {

			log.Println("REST server not run", err)
		}
	}()
}

// Запускает gRPC сервис
func (s *Server) startGRPC() error {
	port := s.options.GrpcPort
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))

	if err != nil {
		log.Printf("Unable to start gRPC service on port %d", port)
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, s.server)

	log.Printf("gRPC service running on port %d", port)

	return grpcServer.Serve(lis)
}

// Запускает REST сервис
func (s *Server) startREST() error {
	port := s.options.RESTPort
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterRouteGuideHandlerFromEndpoint(ctx, gwMux, fmt.Sprintf(":%d", s.options.GrpcPort), opts)

	if err != nil {
		log.Printf("Unable to start REST service on port %d", port)
		return err
	}

	mux.Handle("/", gwMux)

	log.Printf("gRPC service running on port %d", port)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

// Ожидаю завершения всех процессов
func (s *Server) WaitStop() {
	s.wg.Wait()
}
