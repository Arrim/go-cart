package models

import "github.com/arrim/go-cart/pb"

type Cart struct {
	ID       int32             `json:"id"`
	Products []*pb.ProductItem `json:"products"`
}

func (c *Cart) AddProduct(item *pb.Product, quantity int32) {
	for _, p := range c.Products {
		if p.Product.Id == item.Id {
			p.Quantity = p.Quantity + quantity
			p.Product = item
			p.Amount = p.Product.Price * p.Quantity

			return
		}
	}

	c.Products = append(c.Products, &pb.ProductItem{
		Product:  item,
		Quantity: quantity,
		Amount:   item.Price * quantity,
	})
}

func (c *Cart) RemoveProduct(productId int32, quantity int32) {
	for i, p := range c.Products {
		if p.Product.Id == productId {
			if quantity >= p.Quantity {
				c.Products = append(c.Products[:i], c.Products[i+1:]...)
			} else {
				p.Quantity = p.Quantity - quantity
				p.Amount = p.Product.Price * p.Quantity
			}

			return
		}
	}
}
