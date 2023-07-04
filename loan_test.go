package loancalculator_test

import (
	"fmt"
	"testing"
	"time"

	loancalculator "github.com/ThiagoDonadel/loan-calculator"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoanTestSuite struct {
	suite.Suite
}

func (s *LoanTestSuite) TestPriceCalculation() {

	baseDate, _ := time.Parse(time.DateOnly, "2020-10-05")
	params := loancalculator.CalculationParameters{
		InitialValue:   10000.0,
		Method:         loancalculator.FRENCH_PRICE,
		Rate:           5,
		RateBaseMonths: loancalculator.YEARLY,
		Term:           12,
		BaseDate:       baseDate,
	}

	expectedInstallment := 855.57
	expected := []*loancalculator.Value{
		{Number: 0, PaymentDate: baseDate, Installment: 0, Interest: 0, Amortization: 0, Balance: 10000.0},
		{Number: 1, PaymentDate: baseDate.AddDate(0, 1, 0), Installment: expectedInstallment, Interest: 40.74, Amortization: 814.82, Balance: 9185.18},
		{Number: 2, PaymentDate: baseDate.AddDate(0, 2, 0), Installment: expectedInstallment, Interest: 37.42, Amortization: 818.14, Balance: 8367.03},
		{Number: 3, PaymentDate: baseDate.AddDate(0, 3, 0), Installment: expectedInstallment, Interest: 34.09, Amortization: 821.48, Balance: 7545.55},
		{Number: 4, PaymentDate: baseDate.AddDate(0, 4, 0), Installment: expectedInstallment, Interest: 30.74, Amortization: 824.82, Balance: 6720.73},
		{Number: 5, PaymentDate: baseDate.AddDate(0, 5, 0), Installment: expectedInstallment, Interest: 27.38, Amortization: 828.18, Balance: 5892.54},
		{Number: 6, PaymentDate: baseDate.AddDate(0, 6, 0), Installment: expectedInstallment, Interest: 24.01, Amortization: 831.56, Balance: 5060.98},
		{Number: 7, PaymentDate: baseDate.AddDate(0, 7, 0), Installment: expectedInstallment, Interest: 20.62, Amortization: 834.95, Balance: 4226.04},
		{Number: 8, PaymentDate: baseDate.AddDate(0, 8, 0), Installment: expectedInstallment, Interest: 17.22, Amortization: 838.35, Balance: 3387.69},
		{Number: 9, PaymentDate: baseDate.AddDate(0, 9, 0), Installment: expectedInstallment, Interest: 13.80, Amortization: 841.76, Balance: 2545.93},
		{Number: 10, PaymentDate: baseDate.AddDate(0, 10, 0), Installment: expectedInstallment, Interest: 10.37, Amortization: 845.19, Balance: 1700.73},
		{Number: 11, PaymentDate: baseDate.AddDate(0, 11, 0), Installment: expectedInstallment, Interest: 6.93, Amortization: 848.64, Balance: 852.09},
		{Number: 12, PaymentDate: baseDate.AddDate(0, 12, 0), Installment: expectedInstallment, Interest: 3.47, Amortization: 852.09, Balance: 0},
	}

	values, err := loancalculator.Calculate(params)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expected, values)
}

func (s *LoanTestSuite) TestSACCalculation() {

	baseDate, _ := time.Parse(time.DateOnly, "2020-10-05")
	params := loancalculator.CalculationParameters{
		InitialValue:   10000.0,
		Method:         loancalculator.CONSTANT_AMORTIZATION,
		Rate:           5,
		RateBaseMonths: loancalculator.YEARLY,
		Term:           12,
		BaseDate:       baseDate,
	}

	expectedPrincipalValue := 833.33
	expected := []*loancalculator.Value{
		{Number: 0, PaymentDate: baseDate, Installment: 0, Interest: 0, Amortization: 0, Balance: 10000.0},
		{Number: 1, PaymentDate: baseDate.AddDate(0, 1, 0), Installment: 874.07, Interest: 40.74, Amortization: expectedPrincipalValue, Balance: 9166.67},
		{Number: 2, PaymentDate: baseDate.AddDate(0, 2, 0), Installment: 870.68, Interest: 37.35, Amortization: expectedPrincipalValue, Balance: 8333.34},
		{Number: 3, PaymentDate: baseDate.AddDate(0, 3, 0), Installment: 867.28, Interest: 33.95, Amortization: expectedPrincipalValue, Balance: 7500.01},
		{Number: 4, PaymentDate: baseDate.AddDate(0, 4, 0), Installment: 863.89, Interest: 30.56, Amortization: expectedPrincipalValue, Balance: 6666.68},
		{Number: 5, PaymentDate: baseDate.AddDate(0, 5, 0), Installment: 860.49, Interest: 27.16, Amortization: expectedPrincipalValue, Balance: 5833.35},
		{Number: 6, PaymentDate: baseDate.AddDate(0, 6, 0), Installment: 857.10, Interest: 23.77, Amortization: expectedPrincipalValue, Balance: 5000.02},
		{Number: 7, PaymentDate: baseDate.AddDate(0, 7, 0), Installment: 853.70, Interest: 20.37, Amortization: expectedPrincipalValue, Balance: 4166.69},
		{Number: 8, PaymentDate: baseDate.AddDate(0, 8, 0), Installment: 850.31, Interest: 16.98, Amortization: expectedPrincipalValue, Balance: 3333.36},
		{Number: 9, PaymentDate: baseDate.AddDate(0, 9, 0), Installment: 846.91, Interest: 13.58, Amortization: expectedPrincipalValue, Balance: 2500.03},
		{Number: 10, PaymentDate: baseDate.AddDate(0, 10, 0), Installment: 843.52, Interest: 10.19, Amortization: expectedPrincipalValue, Balance: 1666.70},
		{Number: 11, PaymentDate: baseDate.AddDate(0, 11, 0), Installment: 840.12, Interest: 6.79, Amortization: expectedPrincipalValue, Balance: 833.37},
		{Number: 12, PaymentDate: baseDate.AddDate(0, 12, 0), Installment: 836.77, Interest: 3.40, Amortization: 833.37, Balance: 0},
	}

	values, err := loancalculator.Calculate(params)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), expected, values)

}

func TestLoanTestSuite(t *testing.T) {
	suite.Run(t, new(LoanTestSuite))
}

func TestB(t *testing.T) {

	baseDate, _ := time.Parse(time.DateOnly, "2020-10-05")
	params := loancalculator.CalculationParameters{
		InitialValue:   500,
		Method:         loancalculator.FRENCH_PRICE,
		Rate:           1,
		RateBaseMonths: loancalculator.MONTHLY,
		Term:           5,
		BaseDate:       baseDate,
	}

	values, _ := loancalculator.Calculate(params)

	for _, v := range values {
		fmt.Println(v)
	}

}
