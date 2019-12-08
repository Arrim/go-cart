package cart

import (
	"context"
	"errors"
	"github.com/arrim/go-cart/app/models"
	"github.com/arrim/go-cart/app/repository"
	"github.com/arrim/go-cart/pb"
)

type mockCartRepo struct {
}

type Carts []*models.Cart

var db = Carts{
	{
		ID: 1,
		Products: []*pb.ProductItem{
			{
				Product: &pb.Product{
					Id:    1,
					Price: 100,
				},
				Quantity: 1,
				Amount:   100,
			},
		},
	},
}

func SetData(carts Carts) {
	db = carts
}

func (p mockCartRepo) GetByID(ctx context.Context, id int32) (*models.Cart, error) {
	for _, c := range db {
		if c.ID == id {
			return c, nil
		}
	}

	return &models.Cart{}, errors.New("no rows in result set")
}

func (p mockCartRepo) Create(ctx context.Context, model *models.Cart) (int32, error) {
	model.ID = int32(len(db) + 1)
	db = append(db, model)

	return model.ID, nil
}

func (p mockCartRepo) Update(ctx context.Context, model *models.Cart) (int32, error) {
	for i, c := range db {
		if c.ID == model.ID {
			db[i] = model
		}
	}

	return model.ID, nil
}

func (p mockCartRepo) Delete(ctx context.Context, id int32) (bool, error) {
	return false, nil
}

func NewMockCartRepo() repository.CartRepo {
	return mockCartRepo{}
}
