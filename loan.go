package loancalculator

import (
	"math"
	"time"
)

// Calculates loan values according with parameters.
//
// The interest rates are always transformed to a MONTHLY base during the calculation
func Calculate(parameters CalculationParameters) ([]*LoanValue, error) {

	if ok, err := parameters.Validate(); !ok {
		return nil, err
	}

	//set hour, minute, and second to zero
	date, _ := time.Parse(time.DateOnly, parameters.BaseDate.Format(time.DateOnly))
	rate := parameters.Rate / 100

	switch parameters.Method {
	case CONSTANT_AMORTIZATION:
		return calculateConstantAmortization(parameters.InitialValue, rate, parameters.RateBaseMonths, parameters.Term, date)
	case FRENCH_PRICE:
		return calculateFrenchPrice(parameters.InitialValue, rate, parameters.RateBaseMonths, parameters.Term, date)
	default:
		return nil, ErrInvalidMethod
	}
}

// Calculate values following the FRENCH PRICE METHOD rules.
func calculateFrenchPrice(initialValue, rate float64, rateType RateBase, term int, baseDate time.Time) ([]*LoanValue, error) {

	//calculates de equivalent interest MONTHLY rate
	montthlyRate := calculateRate(rate, 1, int(rateType))

	//Calculates the full period rate, using the equivalent rate thba was calculated previously.
	periodRate := calculateRate(montthlyRate, term, int(MONTHLY))

	installmentValue := initialValue * montthlyRate * (periodRate + 1)
	installmentValue = installmentValue / periodRate

	currentDate := baseDate
	finalDate := currentDate.AddDate(0, term, 0)

	payments := []*LoanValue{}
	balance := initialValue

	installmentNumber := 0

	payments = append(payments, &LoanValue{Number: installmentNumber, PaymentDate: currentDate, Balance: balance})

	for ok := true; ok; ok = currentDate.Before(finalDate) {
		currentDate = currentDate.AddDate(0, 1, 0)
		installmentNumber++

		interestValue := balance * montthlyRate
		amortizationValue := installmentValue - interestValue
		balance -= amortizationValue

		if balance < 0 {
			balance = 0
		}

		loanValue := &LoanValue{Number: installmentNumber, PaymentDate: currentDate, Installment: installmentValue, Interest: interestValue, Amortization: amortizationValue, Balance: balance}
		payments = append(payments, loanValue)
	}

	roundValues(payments)

	return payments, nil

}

func calculateConstantAmortization(initialValue, rate float64, rateType RateBase, term int, baseDate time.Time) ([]*LoanValue, error) {

	//calculates de equivalent interest MONTHLY rate
	montthlyRate := calculateRate(rate, 1, int(rateType))

	amortizationValue := initialValue / float64(term)
	amortizationValue = round(amortizationValue)

	currentDate := baseDate
	finalDate := currentDate.AddDate(0, term, 0)

	payments := []*LoanValue{}

	balance := initialValue

	installmentNumber := 0

	payments = append(payments, &LoanValue{Number: installmentNumber, PaymentDate: currentDate, Balance: balance})

	for ok := true; ok; ok = currentDate.Before(finalDate) {
		currentDate = currentDate.AddDate(0, 1, 0)
		installmentNumber++

		if currentDate.Equal(finalDate) {
			amortizationValue = balance
		}

		interestValue := balance * montthlyRate
		installmentValue := amortizationValue + interestValue
		balance -= amortizationValue

		if balance < 0 {
			balance = 0
		}

		loanValue := &LoanValue{Number: installmentNumber, PaymentDate: currentDate, Installment: installmentValue, Interest: interestValue, Amortization: amortizationValue, Balance: balance}
		payments = append(payments, loanValue)

	}

	roundValues(payments)

	return payments, nil

}

func calculateRate(rate float64, period, base int) float64 {
	return math.Pow(1+rate, float64(period)/float64(base)) - 1
}

func roundValues(payments []*LoanValue) {

	for _, payment := range payments {
		payment.Installment = round(payment.Installment)
		payment.Interest = round(payment.Interest)
		payment.Amortization = round(payment.Amortization)
		payment.Balance = round(payment.Balance)
	}

}

func round(value float64) float64 {
	return math.Round(value*100) / 100
}
