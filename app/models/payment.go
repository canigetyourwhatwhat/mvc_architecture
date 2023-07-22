package entity

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Payment struct {
	ID        int           `db:"id"`
	UserId    int           `db:"userId"`
	Amount    float32       `db:"amount"`
	Method    PaymentMethod `db:"method"`
	CreatedAt time.Time     `db:"createdAt"`
	UpdatedAt time.Time     `db:"updatedAt"`
}

type MakePaymentInput struct {
	PaymentMethod int `json:"paymentMethod"`
}

type PaymentMethod int

const (
	Card PaymentMethod = iota
	Cash
)

func (p *Payment) CreatePayment(db *sqlx.DB) (int, error) {
	query := `INSERT INTO payments (amount, userId, method) VALUES (:amount, :userId, :method)`
	result, err := db.NamedExec(query, *p)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (p *Payment) GetPaymentById(db *sqlx.DB, id int) error {
	err := db.Get(p, "select * from payments where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
