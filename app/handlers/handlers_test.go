package handlers

import (
	"context"
	"github.com/arrim/go-cart/app/repository"
	cartMock "github.com/arrim/go-cart/mock"
	"github.com/arrim/go-cart/pb"
	"reflect"
	"testing"
)

func TestNewRouterGuideServer(t *testing.T) {
	type args struct {
		r *repository.Repository
	}

	tests := []struct {
		name string
		args args
		want *RouteGuideServer
	}{
		{
			name: "",
			args: args{
				r: repository.NewRepository(cartMock.NewMockCartRepo()),
			},
			want: &RouteGuideServer{
				repository.NewRepository(cartMock.NewMockCartRepo()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouterGuideServer(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouterGuideServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouteGuideServer_Add(t *testing.T) {
	type fields struct {
		repository *repository.Repository
	}

	type args struct {
		ctx     context.Context
		request *pb.AddRequest
	}

	cartMock.SetData(cartMock.Carts{})

	repo := repository.NewRepository(cartMock.NewMockCartRepo())

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.StatusMessage
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				repository: repo,
			},
			args: args{
				ctx: context.Background(),
				request: &pb.AddRequest{
					Id: 1,
					Product: &pb.Product{
						Id:    1,
						Price: 1000,
					},
					Quantity: 1,
				},
			},
			want: &pb.StatusMessage{
				Ok:      true,
				Code:    0,
				Message: "ok",
				Id:      1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RouteGuideServer{
				repository: tt.fields.repository,
			}
			got, err := s.Add(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouteGuideServer_Delete(t *testing.T) {
	type fields struct {
		repository *repository.Repository
	}
	type args struct {
		ctx     context.Context
		request *pb.DeleteRequest
	}

	repo := repository.NewRepository(cartMock.NewMockCartRepo())

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.StatusMessage
		wantErr bool
	}{
		{
			name: "Delete product",
			fields: fields{
				repository: repo,
			},
			args: args{
				ctx: context.Background(),
				request: &pb.DeleteRequest{
					Id:        1,
					ProductId: 1,
					Quantity:  1,
				},
			},
			want: &pb.StatusMessage{
				Ok:      true,
				Code:    0,
				Message: "ok",
			},
			wantErr: false,
		},
		{
			name: "Delete product negative",
			fields: fields{
				repository: repo,
			},
			args: args{
				ctx: context.Background(),
				request: &pb.DeleteRequest{
					Id:        2,
					ProductId: 1,
					Quantity:  1,
				},
			},
			want: &pb.StatusMessage{
				Ok:      false,
				Code:    0,
				Message: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RouteGuideServer{
				repository: tt.fields.repository,
			}
			got, err := s.Delete(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouteGuideServer_Get(t *testing.T) {
	type fields struct {
		repository *repository.Repository
	}
	type args struct {
		ctx     context.Context
		request *pb.Cart
	}

	cartMock.SetData(cartMock.Carts{
		{
			ID: 1,
			Products: []*pb.ProductItem{
				{
					Product: &pb.Product{
						Id:    1,
						Price: 50,
					},
					Quantity: 2,
					Amount:   100,
				},
				{
					Product: &pb.Product{
						Id:    2,
						Price: 100,
					},
					Quantity: 4,
					Amount:   400,
				},
			},
		},
	})

	repo := repository.NewRepository(cartMock.NewMockCartRepo())

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetResponse
		wantErr bool
	}{
		{
			name: "Get Cart",
			fields: fields{
				repository: repo,
			},
			args: args{
				ctx: context.Background(),
				request: &pb.Cart{
					Id: 1,
				},
			},
			want: &pb.GetResponse{
				Id: 1,
				Products: []*pb.ProductItem{
					{
						Product: &pb.Product{
							Id:    1,
							Price: 50,
						},
						Quantity: 2,
						Amount:   100,
					},
					{
						Product: &pb.Product{
							Id:    2,
							Price: 100,
						},
						Quantity: 4,
						Amount:   400,
					},
				},
				Quantity: 6,
				Amount:   500,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RouteGuideServer{
				repository: tt.fields.repository,
			}
			got, err := s.Get(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRouteGuideServer_Health(t *testing.T) {
	type fields struct {
		repository *repository.Repository
	}
	type args struct {
		ctx   context.Context
		empty *pb.Empty
	}
	repo := repository.NewRepository(cartMock.NewMockCartRepo())

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.HealthMessage
		wantErr bool
	}{
		{
			name: "Health",
			fields: fields{
				repository: repo,
			},
			args: args{
				ctx:   context.Background(),
				empty: &pb.Empty{},
			},
			want: &pb.HealthMessage{
				Message: "ok",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RouteGuideServer{
				repository: tt.fields.repository,
			}
			got, err := s.Health(tt.args.ctx, tt.args.empty)
			if (err != nil) != tt.wantErr {
				t.Errorf("Health() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Health() got = %v, want %v", got, tt.want)
			}
		})
	}
}
