package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type Payment struct {
	ID        string
	Amount    decimal.Decimal
	Method    paymentMethod
	CreatedAt time.Time
	UpdatedAt time.Time
}

type paymentMethod int

const (
	Card paymentMethod = iota
	Cash
)
