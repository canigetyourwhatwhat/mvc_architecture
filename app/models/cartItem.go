package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type CartItem struct {
	ID         string
	Product    Product
	Qty        int
	TotalPrice decimal.Decimal
	TaxPrice   decimal.Decimal
	NetPrice   decimal.Decimal
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
