package loancalculator_test

import (
	"errors"
	"testing"
	"time"

	loancalculator "github.com/ThiagoDonadel/loan-calculator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RateBaseTestSuite struct {
	suite.Suite
}

func (s *RateBaseTestSuite) TestValidRateBase() {

	ratebase := loancalculator.RateBase(loancalculator.YEARLY)

	isValid := ratebase.IsValid()

	assert.True(s.T(), isValid)
}

func (s *RateBaseTestSuite) TestInvalidRateBase() {

	ratebase := loancalculator.RateBase(15)

	isValid := ratebase.IsValid()

	assert.False(s.T(), isValid)
}

func TestRateBaseTestSuite(t *testing.T) {
	suite.Run(t, new(RateBaseTestSuite))
}

type CalculationMethodTestSuite struct {
	suite.Suite
}

func (s *CalculationMethodTestSuite) TestValidCalculationMethod() {

	method := loancalculator.CalculationMethod(loancalculator.CONSTANT_AMORTIZATION)

	isValid := method.IsValid()

	assert.True(s.T(), isValid)
}

func (s *CalculationMethodTestSuite) TestInvalidCalculationMethod() {

	method := loancalculator.CalculationMethod(88)

	isValid := method.IsValid()

	assert.False(s.T(), isValid)
}

func TestCalculationMethodTestSuite(t *testing.T) {
	suite.Run(t, new(CalculationMethodTestSuite))
}

type CalculationParametersTestSuite struct {
	suite.Suite
}

func (s *CalculationParametersTestSuite) TestValidCalculationParameters() {

	params := loancalculator.CalculationParameters{
		Method:         loancalculator.FRENCH_PRICE,
		InitialValue:   10000,
		Rate:           10,
		RateBaseMonths: loancalculator.MONTHLY,
		Term:           12,
		BaseDate:       time.Now(),
	}

	isValid, err := params.Validate()

	assert.True(s.T(), isValid)
	assert.Nil(s.T(), err)
}

func (s *CalculationParametersTestSuite) TestInvalidCalculationParameters() {

	params := loancalculator.CalculationParameters{
		Method:         loancalculator.CalculationMethod(5),
		InitialValue:   -1000,
		Rate:           -10,
		RateBaseMonths: loancalculator.RateBase(22),
	}

	isValid, err := params.Validate()

	assert.False(s.T(), isValid)
	assert.True(s.T(), errors.Is(err, loancalculator.ErrInvalidBaseDate))
	assert.True(s.T(), errors.Is(err, loancalculator.ErrInvalidInitalValue))
	assert.True(s.T(), errors.Is(err, loancalculator.ErrInvalidMethod))
	assert.True(s.T(), errors.Is(err, loancalculator.ErrInvalidPeriod))
	assert.True(s.T(), errors.Is(err, loancalculator.ErrInvalidRate))
	assert.True(s.T(), errors.Is(err, loancalculator.ErrInvalidRateBaseMonths))
}

func TestCalculationParametersTestSuite(t *testing.T) {
	suite.Run(t, new(CalculationParametersTestSuite))
}
