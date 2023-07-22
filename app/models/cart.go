package entity

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type Cart struct {
	ID         int        `db:"id"`
	UserId     int        `db:"userId"`
	TotalPrice float32    `db:"totalPrice"`
	TaxPrice   float32    `db:"taxPrice"`
	NetPrice   float32    `db:"netPrice"`
	Status     CartStatus `db:"status"`
	CartItems  []CartItem
	CreatedAt  time.Time `db:"createdAt"`
	UpdatedAt  time.Time `db:"updatedAt"`
}

type CartStatus int

const (
	InProgress CartStatus = iota
	Completed
)

func (c *Cart) GetCartById(db *sqlx.DB, id int) error {
	err := db.Get(c, "select * from carts where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cart) GetCartByUserId(db *sqlx.DB, userId int) ([]Cart, error) {
	var carts []Cart
	err := db.Select(&carts, "select * from carts where userId = ?", userId)
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (c *Cart) GetInProgressCartByUserId(db *sqlx.DB, userId int) (*Cart, error) {
	var cart Cart
	err := db.Get(&cart, "select * from carts where userId = ? and status = ?", userId, InProgress)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (c *Cart) CreateCart(db *sqlx.DB, userId int) error {
	cart := Cart{UserId: userId}
	query := `INSERT INTO carts (userId) VALUES (:userId)`
	_, err := db.NamedExec(query, cart)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cart) UpdateCart(db *sqlx.DB) error {
	query := `UPDATE carts set netPrice = :netPrice, taxPrice = :taxPrice, totalPrice = :totalPrice, status = :status where id = :id`
	_, err := db.NamedExec(query, c)
	if err != nil {
		fmt.Println("\n\n hit ")
		return err
	}
	return nil
}
