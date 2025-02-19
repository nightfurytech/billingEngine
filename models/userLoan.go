package models

import "time"

type Loan struct {
	UserID        int       `json:"userId"`
	LoanStartDate time.Time `json:"loanStartDate"`
	LoanEndDate   time.Time `json:"loanEndDate"`
	AmountTaken   float64   `json:"amountTaken"`
	AmountPaid    float64   `json:"amountPaid"`
}
