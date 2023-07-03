package entity

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type CartItem struct {
	ID          string    `db:"id"`
	ProductCode string    `db:"productCode"`
	CartId      int       `db:"cartId"`
	Quantity    int       `db:"quantity"`
	TotalPrice  float32   `db:"totalPrice"`
	TaxPrice    float32   `db:"taxPrice"`
	NetPrice    float32   `db:"netPrice"`
	CreatedAt   time.Time `db:"createdAt"`
	UpdatedAt   time.Time `db:"updatedAt"`
}

type AddCartItemRequest struct {
	SessionId   string `json:"sessionId"`
	ProductCode string `json:"productCode"`
	Quantity    int    `json:"quantity"`
}

type DeleteCartItemRequest struct {
	SessionId   string `json:"sessionId"`
	ProductCode string `json:"productCode"`
}

func (ci *CartItem) CreateItemInCart(db *sqlx.DB) error {
	query := `INSERT INTO cartItems (productCode, cartId, quantity, totalPrice, taxPrice, netPrice) VALUES (:productCode, :cartId, :quantity, :totalPrice, :taxPrice, :netPrice) ON DUPLICATE KEY UPDATE quantity = :quantity, totalPrice = :totalPrice, taxPrice = :taxPrice, netPrice = :netPrice`
	_, err := db.NamedExec(query, *ci)
	if err != nil {
		return err
	}
	return nil
}

func (ci *CartItem) DeleteItemInCart(db *sqlx.DB) error {
	query := `delete from cartItems where id = :id`
	_, err := db.NamedExec(query, *ci)
	if err != nil {
		return err
	}
	return nil
}

func (ci *CartItem) GetCartItemByProductIdAndCartId(db *sqlx.DB, productCode string, cartId int) (*CartItem, error) {
	var cartItem CartItem
	err := db.Get(&cartItem, "select * from cartItems where productCode = ? and cartId = ?", productCode, cartId)
	if err != nil {
		return nil, err
	}
	return &cartItem, nil
}
