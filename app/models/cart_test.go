package models

import (
	"github.com/arrim/go-cart/pb"
	"math/rand"
	"testing"
)

func TestAddNewProduct(t *testing.T) {
	c := Cart{}
	q := rand.Int31n(10)
	price := rand.Int31()

	c.AddProduct(
		&pb.Product{
			Id:    0,
			Name:  "phone",
			Price: price,
		},
		q,
	)

	if len(c.Products) != 1 {
		t.Fatal("Список товаров пуст")
	}

	if c.Products[0].Quantity != q {
		t.Fatal("Количество товара не соответствует")
	}

	if c.Products[0].Amount != price*q {
		t.Fatal("Стоимость товаров не соответствует")
	}
}

func TestAddProduct(t *testing.T) {
	c := Cart{}
	q := rand.Int31n(10)
	q2 := rand.Int31n(10)
	price := rand.Int31()

	c.AddProduct(
		&pb.Product{
			Id:    0,
			Name:  "phone",
			Price: price,
		},
		q,
	)

	c.AddProduct(
		&pb.Product{
			Id:    0,
			Name:  "phone",
			Price: price,
		},
		q2,
	)

	if len(c.Products) != 1 {
		t.Fatal("Список товаров пуст")
	}

	if c.Products[0].Quantity != q+q2 {
		t.Fatal("Количество товара не соответствует")
	}

	if c.Products[0].Amount != price*(q+q2) {
		t.Fatal("Стоимость товаров не соответствует")
	}
}

func TestRemoveProduct(t *testing.T) {
	c := Cart{}
	q := rand.Int31n(10)
	d := rand.Int31n(q)
	id := rand.Int31()
	price := rand.Int31()

	c.AddProduct(
		&pb.Product{
			Id:    id,
			Price: price,
		},
		q,
	)

	c.RemoveProduct(id, d)

	if q-d > 0 {
		if c.Products[0].Quantity != q-d {
			t.Fatal("Количество товара не соответствует")
		}

		if c.Products[0].Amount != price*(q-d) {
			t.Fatal("Стоимость товаров не соответствует")
		}
	} else {
		if len(c.Products) != 0 {
			t.Fatal("Список товаров не пуст")
		}
	}

}
