package entity

type Order struct {
	User    User
	Cart    Cart
	Payment Payment
}
