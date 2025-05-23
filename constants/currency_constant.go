package constants

type Currency string

const (
	EUR Currency = "Euro"
	INR Currency = "Indian Rupees"
	USD Currency = "U.S. Dollars"
)

var Rates = map[Currency]float64{
	INR: 1,
	EUR: 0.02,
	USD: 0.05,
}

func (c Currency) IsValidCurrency() bool {
	switch c {
	case EUR:
		return true
	case INR:
		return true
	case USD:
		return true
	default:
		return false
	}
}

func (c Currency) CurrencyConverter(from, to Currency, val float64) float64 {

	if from == INR {
		return val * Rates[to]
	} else {
		inrValue := val / Rates[from]

		return inrValue * Rates[to]
	}

}
