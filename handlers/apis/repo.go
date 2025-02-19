package apis

import (
	"billingEngine/models"
	"go.uber.org/zap"
	"sync"
)

type ApiHandler struct {
	logger      *zap.Logger
	userLoanMap *sync.Map
}

func NewApiHandler(logger *zap.Logger, userLoans []models.Loan) *ApiHandler {
	userLoanMap := &sync.Map{}
	for _, loan := range userLoans {
		userLoanMap.Store(loan.UserID, loan)
	}
	return &ApiHandler{logger: logger, userLoanMap: userLoanMap}
}
