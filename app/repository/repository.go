package repository

import (
	"context"
	"github.com/arrim/go-cart/app/models"
)

type Repository struct {
	Cart CartRepo
}

type CartRepo interface {
	GetByID(ctx context.Context, id int32) (*models.Cart, error)
	Create(ctx context.Context, model *models.Cart) (int32, error)
	Update(ctx context.Context, model *models.Cart) (int32, error)
	Delete(ctx context.Context, id int32) (bool, error)
}

func NewRepository(cart CartRepo) *Repository {
	return &Repository{
		Cart: cart,
	}
}
