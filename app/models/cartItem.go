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

type CartItemRequest struct {
	ProductCode string `json:"productCode"`
	Quantity    int    `json:"quantity"`
}

type UpdateCartItemRequest struct {
	Records []CartItemRequest `json:"records"`
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

func (ci *CartItem) GetCarItemsByCartId(db *sqlx.DB, cartId int) (cartItems []CartItem, err error) {
	err = db.Select(&cartItems, "select * from cartItems where cartId = ?", cartId)
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}
