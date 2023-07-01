package entity

import (
	"time"
)

type CartItem struct {
	ID         string    `db:"id"`
	ProductId  string    `db:"productId"`
	CartId     string    `db:"cartId"`
	Qty        int       `db:"qty"`
	TotalPrice float32   `db:"totalPrice"`
	TaxPrice   float32   `db:"TaxPrice"`
	NetPrice   float32   `db:"NetPrice"`
	CreatedAt  time.Time `db:"createdAt"`
	UpdatedAt  time.Time `db:"updatedAt"`
}
