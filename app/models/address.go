package entity

import "time"

type Address struct {
	CountryID  string
	ProvinceID string
	CityID     string
	Address1   string
	Address2   string
	PostCode   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
