package entity

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Cart struct {
	ID         int        `db:"id"`
	UserId     string     `db:"userId"`
	TotalPrice float32    `db:"totalPrice"`
	TaxPrice   float32    `db:"taxPrice"`
	NetPrice   float32    `db:"netPrice"`
	Status     CartStatus `db:"status"`
	CreatedAt  time.Time  `db:"createdAt"`
	UpdatedAt  time.Time  `db:"updatedAt"`
}

type CartStatus int

const (
	InProgress CartStatus = iota
	Completed
)

func convertCartStatus(status CartStatus) string {
	switch status {
	case InProgress:
		return "0"
	case Completed:
		return "1"
	}
	return ""
}

func (c *Cart) GetInProgressCartByUserId(db *sqlx.DB, userId string) (*Cart, error) {
	var cart Cart
	err := db.Get(&cart, "select * from carts where userId = ? and status = ?", userId, convertCartStatus(InProgress))
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (c *Cart) CreateCart(db *sqlx.DB, userId string) error {
	cart := Cart{UserId: userId}
	query := `INSERT INTO carts (userId) VALUES (:userId)`
	_, err := db.NamedExec(query, cart)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cart) UpdateCart(db *sqlx.DB) error {
	query := `UPDATE carts set netPrice = :netPrice, taxPrice = :taxPrice, totalPrice = :totalPrice`
	_, err := db.NamedExec(query, c)
	if err != nil {
		return err
	}
	return nil
}
