package util

const (
	USD = "USD"
	EUR = "EUR"
	VND = "VND"
)

func AcceptedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, VND:
		return true
	}
	return false
}
