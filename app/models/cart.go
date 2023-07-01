package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type Cart struct {
	ID         string
	CartItems  []CartItem
	TotalPrice decimal.Decimal
	TaxPrice   decimal.Decimal
	NetPrice   decimal.Decimal
	Status     CartStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CartStatus int

const (
	InProgress CartStatus = iota
	Completed
)
