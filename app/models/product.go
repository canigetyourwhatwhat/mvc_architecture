package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type Product struct {
	ID               string
	Name             string
	Price            decimal.Decimal
	Stock            int
	ShortDescription string
	LongDescription  string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
