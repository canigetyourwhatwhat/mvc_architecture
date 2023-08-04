package entity

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Order struct {
	ID        int       `db:"id"`
	UserId    int       `db:"userId"`
	CartId    int       `db:"cartId"`
	PaymentId int       `db:"paymentId"`
	CreatedAt time.Time `db:"createdAt"`
}

type OrderResponse struct {
	Payment Payment `json:"payment"`
	Cart    Cart    `json:"cart"`
}

func (o *Order) CreateOrder(db *sqlx.DB) error {
	query := `INSERT INTO orders (userId, cartId, paymentId) VALUES (:userId, :cartId, :paymentId)`
	_, err := db.NamedExec(query, *o)
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) GetOrder(db *sqlx.DB, id int) error {
	err := db.Get(o, "select * from orders where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) ListOrdersByUserId(db *sqlx.DB, userId string) ([]*Order, error) {
	var orders []*Order
	err := db.Select(&orders, "select * from orders where userId = ? order by createdAt desc", userId)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
