package util

// Constants for all supported currencies
const (
	USD   = "USD"
	EUR   = "EUR"
	TENGE = "TENGE"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, TENGE:
		return true
	}
	return false
}
