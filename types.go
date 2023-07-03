package loancalculator

import (
	"errors"
	"time"
)

// Struct that holds the calculated values
type LoanValue struct {
	Number       int       // installment number
	PaymentDate  time.Time // date of the installment
	Installment  float64   // installment value paid.
	Interest     float64   // interest paid on the date
	Amortization float64   // the amortization value paid on the date
	Balance      float64   // remaining loanaed ammount
}

// Struct that holds the parameters that will be used to calculate
type CalculationParameters struct {
	Method         CalculationMethod // method to be used to calculate
	InitialValue   float64           // value loaned, must be > 0
	Rate           float64           // interest Rate Value (percentual value), must be >= 0
	RateBaseMonths RateBase          // number of months that the interest occurs (e.g. a interest that occurs Yearly has a period of 12 months)
	Term           int               // number of months that the calculation will run, must be > 0
	BaseDate       time.Time         // date that the calculation starts
}

// Validate if the informed parameters are valid
func (c *CalculationParameters) Validate() (bool, error) {

	var err error = nil

	if c.InitialValue <= 0 {
		err = errors.Join(err, ErrInvalidInitalValue)
	}

	if c.Rate < 0 {
		err = errors.Join(err, ErrInvalidRate)
	}

	if c.Term <= 0 {
		err = errors.Join(err, ErrInvalidPeriod)
	}

	if c.BaseDate.IsZero() {
		err = errors.Join(err, ErrInvalidBaseDate)
	}

	if !c.RateBaseMonths.IsValid() {
		err = errors.Join(err, ErrInvalidRateBaseMonths)
	}

	if !c.Method.IsValid() {
		err = errors.Join(err, ErrInvalidMethod)
	}

	return err == nil, err
}

type RateBase int

const (
	MONTHLY      RateBase = 1
	QUARTERLY    RateBase = 4
	SEMIANNUALLY RateBase = 6
	YEARLY       RateBase = 12
)

// Validate if the value is valid for the calculation
func (r RateBase) IsValid() bool {
	switch r {
	case MONTHLY, QUARTERLY, SEMIANNUALLY, YEARLY:
		return true
	default:
		return false
	}
}

type CalculationMethod int

const (
	CONSTANT_AMORTIZATION CalculationMethod = iota
	FRENCH_PRICE
)

// Validate if the value is valid for the calculation
func (c CalculationMethod) IsValid() bool {
	switch c {
	case CONSTANT_AMORTIZATION, FRENCH_PRICE:
		return true
	default:
		return false
	}
}

var (
	ErrInvalidInitalValue    = errors.New("parameter InitialValue has a invalid value")
	ErrInvalidRate           = errors.New("parameter Rate has a invalid value")
	ErrInvalidPeriod         = errors.New("parameter Period has a invalid value")
	ErrInvalidBaseDate       = errors.New("parameter BaseDate has a invalid value")
	ErrInvalidRateBaseMonths = errors.New("parameter RateBaseMonths has a invalid value")
	ErrInvalidMethod         = errors.New("parameter Method has a invalid value")
)
