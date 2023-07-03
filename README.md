
# **Loan Calculator** :moneybag:

## **About The Project**

This is a study project/library that calculates the values of a loan. A previous knowledge about finances, loans and mathematics are necessary to fully understand this application

## **Getting Started**

### **Instaling**

Use go get to install the latest version of the library.

```sh
github.com/ThiagoDonadel/funding-calculator
```

### **Usage**

You can do a simple calculation calling the "Calculate" method

```go
loancalculator.Calculate(params)
```

The method will return a slice of calculated values.

### **Calculation Parameters**

Parameters are defined by the **CalculationParameters** structure

```go
type CalculationParameters struct {
	Method         CalculationMethod //method to be used to calculate
	InitialValue   float64           //value loaned, must be > 0
	Rate           float64           //interest Rate Value (percentual value), must be >= 0
	RateBaseMonths RateBase          //number of months that the interest occurs (e.g. a interest that occurs Yearly has a period of 12 months)
	Term           int               //number of months that the calculation will run, must be > 0
	BaseDate       time.Time         //date that the calculation starts
}
```
## **Calculation Rules**

This section will explain some rules used in the calculations.

### **Calculation Method**

The library has two calculations  methods: 

**Constant Amortization:**  In this method we have a constant amortization value being paid every month

E.g.:

|             | Installment | Interest | Amortization | Balance |
| ----------- | ----------- | -------- | ------------ | ------- |
| 0           |        0.00 |     0.00 |         0.00 |  500.00 |
| 1           |      105.00 |     5.00 |       100.00 |  400.00 |
| 2           |      104.00 |     4.00 |       100.00 |  300.00 |
| 3           |      103.00 |     3.00 |       100.00 |  200.00 |
| 4           |      102.00 |     2.00 |       100.00 |  100.00 |
| 5           |      101.00 |     1.00 |       100.00 |    0.00 |


**French Price:** - In this method we have a constant installment value being paid every month. 

E.g.:

|             | Installment | Interest | Amortization | Balance |
| ----------- | ----------- | -------- | ------------ | ------- |
| 0           |        0.00 |     0.00 |         0.00 |  500.00 |
| 1           |      103.02 |     5.00 |        98.02 |  401.98 |
| 2           |      103.02 |     4.02 |        99.00 |  302.98 |
| 3           |      103.02 |     3.03 |        99.99 |  202.99 |
| 4           |      103.02 |     2.03 |       100.99 |  102.00 |
| 5           |      103.02 |     1.02 |       102.00 |    0.00 |

### **Interest Rates**

All calculations are calculated **MONTHLY**, because of that the interest rate is always transformed to monthly rate using the equivalency calculation.

All interest values are calculated using the compound rule.

### **Installment Date Generator**

The loan payments dates are automatically generated, starting at baseDate + 1 month, and sequentially adding a month until reaching baseDate + term 