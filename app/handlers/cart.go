package handlers

import (
	"context"
	"github.com/arrim/go-cart/app/models"
	"github.com/arrim/go-cart/pb"
)

func (s *RouteGuideServer) Add(ctx context.Context, request *pb.AddRequest) (*pb.StatusMessage, error) {
	var model *models.Cart
	var err error

	model = &models.Cart{}

	if request.Id != 0 {
		model, err = s.repository.Cart.GetByID(ctx, request.Id)

		if err != nil && err.Error() == "no rows in result set" {
			model.ID = request.Id
		} else if err != nil {
			return &pb.StatusMessage{
				Ok: false,
			}, err
		}
	}

	model.AddProduct(request.Product, request.Quantity)

	if model.ID == 0 {
		id, err := s.repository.Cart.Create(ctx, model)
		model.ID = id

		if err != nil {
			return &pb.StatusMessage{
				Ok: false,
			}, err
		}
	} else {
		_, err := s.repository.Cart.Update(ctx, model)

		if err != nil {
			return &pb.StatusMessage{
				Ok: false,
			}, err
		}
	}

	return &pb.StatusMessage{
		Ok:      true,
		Code:    0,
		Message: "ok",
		Id:      model.ID,
	}, nil
}

func (s *RouteGuideServer) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.StatusMessage, error) {
	model, err := s.repository.Cart.GetByID(ctx, request.Id)

	if err != nil {
		return &pb.StatusMessage{
			Ok: false,
		}, err
	}

	model.RemoveProduct(request.ProductId, request.Quantity)

	_, err = s.repository.Cart.Update(ctx, model)

	if err != nil {
		return &pb.StatusMessage{
			Ok: false,
		}, err
	}

	return &pb.StatusMessage{
		Ok:      true,
		Code:    0,
		Message: "ok",
	}, nil
}

func (s *RouteGuideServer) Get(ctx context.Context, request *pb.Cart) (*pb.GetResponse, error) {
	var (
		quantity int32
		amount   int32
	)

	quantity = 0
	amount = 0

	model, err := s.repository.Cart.GetByID(ctx, request.Id)

	if err != nil {
		return &pb.GetResponse{}, err
	}

	for _, p := range model.Products {
		quantity = quantity + p.Quantity
		amount = amount + p.Amount
	}

	return &pb.GetResponse{
		Id:       model.ID,
		Products: model.Products,
		Quantity: quantity,
		Amount:   amount,
	}, nil
}
