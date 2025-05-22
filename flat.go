package gofinancial

import "github.com/shopspring/decimal"

// Flat implements financial methods for facilitating a loan use case, following a flat rate of interest.
type Flat struct{}

// GetPrincipal returns principal amount contribution in a given period towards a loan, depending on config.
func (f *Flat) GetPrincipal(config Config, _ int64) decimal.Decimal {
	dPeriod := decimal.NewFromInt(config.periods)
	minusOne := decimal.NewFromInt(-1)
	return config.AmountBorrowed.Div(dPeriod).Mul(minusOne)
}

// GetInterest returns interest amount contribution in a given period towards a loan, depending on config.
func (f *Flat) GetInterest(config Config, period int64) decimal.Decimal {
	minusOne := decimal.NewFromInt(-1)
	paidPrincipal := f.GetPrincipal(config, 0).Mul(decimal.NewFromInt(period).Add(minusOne))
	// return config.getInterestRatePerPeriodInDecimal().Mul(config.AmountBorrowed).Mul(minusOne) // original BS
	return config.getInterestRatePerPeriodInDecimal().Mul(config.AmountBorrowed.Add(paidPrincipal)).Mul(minusOne) // AmountBorrowed is negative, hence .Add()
}

// GetPayment returns the periodic payment to be done for a loan depending on config.
func (f *Flat) GetPayment(config Config, period int64) decimal.Decimal {
	// dPeriod := decimal.NewFromInt(config.periods)
	// minusOne := decimal.NewFromInt(-1)
	// totalInterest := config.getInterestRatePerPeriodInDecimal().Mul(dPeriod).Mul(config.AmountBorrowed)
	// Payment := totalInterest.Add(config.AmountBorrowed).Mul(minusOne).Div(dPeriod)
	interest := f.GetInterest(config, period)
	principal := f.GetPrincipal(config, period)
	return principal.Add(interest)
}
