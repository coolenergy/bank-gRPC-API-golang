package util

const (
	USD = "USD"
	EUR = "EUR"
	GEL = "GEL"
)

func ISSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, GEL:
		return true
	}
	return false
}
