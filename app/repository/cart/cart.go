package cart

import (
	"context"
	"github.com/arrim/go-cart/app/models"
	"github.com/arrim/go-cart/app/repository"
	"github.com/arrim/go-cart/pb"

	"github.com/jackc/pgx/v4"
)

type postgresCartRepo struct {
	Conn *pgx.Conn
}

func (p postgresCartRepo) GetByID(ctx context.Context, id int32) (*models.Cart, error) {
	var products []*pb.ProductItem

	query := "Select id, products from carts where id=$1"

	err := p.Conn.QueryRow(ctx, query, id).Scan(&id, &products)

	if err != nil {
		return &models.Cart{}, err
	}

	return &models.Cart{
		ID:       id,
		Products: products,
	}, nil
}

func (p postgresCartRepo) Create(ctx context.Context, model *models.Cart) (int32, error) {
	var id int32
	query := "insert into carts(products) values ($1) returning id"

	rows, err := p.Conn.Query(ctx, query, model.Products)

	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id)

		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (p postgresCartRepo) Update(ctx context.Context, model *models.Cart) (int32, error) {
	query := "update carts set products=$1 where id=$2 returning id"

	rows, err := p.Conn.Query(ctx, query, model.Products, model.ID)

	if err != nil {
		return 0, err
	}

	defer rows.Close()

	return model.ID, nil
}

func (p postgresCartRepo) Delete(ctx context.Context, id int32) (bool, error) {
	return false, nil
}

func NewPostgresCartRepo(Conn *pgx.Conn) repository.CartRepo {
	return postgresCartRepo{
		Conn: Conn,
	}
}
