package gofinancial

import "github.com/shopspring/decimal"

// Financial interface defines the methods to be over ridden for different financial use cases.
type Financial interface {
	GetPrincipal(config Config, period int64) decimal.Decimal
	GetInterest(config Config, period int64) decimal.Decimal
	GetPayment(config Config, period int64) decimal.Decimal // updated for correct FLAT payment calculation
	// GetPayment(config Config) decimal.Decimal // original
}
