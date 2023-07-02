package entity

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type CartItem struct {
	ID         string    `db:"id"`
	ProductId  string    `db:"productId"`
	CartId     int       `db:"cartId"`
	Quantity   int       `db:"quantity"`
	TotalPrice float32   `db:"totalPrice"`
	TaxPrice   float32   `db:"taxPrice"`
	NetPrice   float32   `db:"netPrice"`
	CreatedAt  time.Time `db:"createdAt"`
	UpdatedAt  time.Time `db:"updatedAt"`
}

type AddCartItemRequest struct {
	SessionId   string `json:"sessionId"`
	ProductCode string `json:"productCode"`
	Quantity    int    `json:"quantity"`
}

func (ci *CartItem) CreateItemInCart(db *sqlx.DB) error {
	query := `INSERT INTO cartItems (productId, cartId, quantity, totalPrice, taxPrice, netPrice) VALUES (:productId, :cartId, :quantity, :totalPrice, :taxPrice, :netPrice) ON DUPLICATE KEY UPDATE quantity = :quantity, totalPrice = :totalPrice, taxPrice = :taxPrice, netPrice = :netPrice`
	_, err := db.NamedExec(query, *ci)
	if err != nil {
		return err
	}
	return nil
}
