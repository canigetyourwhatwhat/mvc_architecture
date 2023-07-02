package entity

const (
	TaxPercent = 10
)

func GetTaxPercent() float32 {
	return TaxPercent / 100.0
}
