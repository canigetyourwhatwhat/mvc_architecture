package entity

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Product struct {
	ID               string    `db:"id"`
	Code             string    `db:"code"`
	Name             string    `db:"name"`
	Price            float32   `db:"price"`
	Stock            int       `db:"stock"`
	ShortDescription string    `db:"short_description"`
	LongDescription  string    `db:"long_description"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

func (p *Product) GetProducts(db *sqlx.DB, perPage int, page int) ([]Product, int, error) {
	var products []Product
	var count int

	err := db.Get(&count, "select count(id) from Products")
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * perPage

	err = db.Select(&products, "select * from Products order by created_at desc limit ? offset ?", perPage, offset)

	if err != nil {
		return nil, 0, err
	}

	return products, count, nil
}

func (p *Product) GetProductByCode(db *sqlx.DB, code string) (*Product, error) {
	var product Product

	err := db.Get(&product, "select * from Products where code = ?", code)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
