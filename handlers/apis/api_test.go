package apis

import (
	"billingEngine/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func newMockApiHandler() *ApiHandler {
	return &ApiHandler{
		userLoanMap: &sync.Map{},
	}
}

func TestGetOutStandingAmount(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := newMockApiHandler()
	handler.userLoanMap.Store(1, models.Loan{
		AmountTaken: 1000,
		AmountPaid:  200,
	})

	_, _ = http.NewRequest("GET", "/outstanding/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "userId", Value: "1"}}

	handler.GetOutStandingAmount(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestIsDelinquent(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := newMockApiHandler()
	handler.userLoanMap.Store(1, models.Loan{
		AmountTaken:   1000,
		AmountPaid:    500,
		LoanStartDate: time.Now().AddDate(0, 0, -21),
		LoanEndDate:   time.Now().AddDate(0, 1, 0),
	})

	_, _ = http.NewRequest("GET", "/isDelinquent/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "userId", Value: "1"}}

	handler.IsDelinquent(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMakePayment(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := newMockApiHandler()
	handler.userLoanMap.Store(1, models.Loan{
		LoanStartDate: time.Now(),
		LoanEndDate:   time.Now().AddDate(0, 0, 350),
		AmountTaken:   5000000,
		AmountPaid:    7 * 110000,
	})

	req, _ := http.NewRequest("POST", "/payment/1?amount=110000", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "userId", Value: "1"}}
	c.Request = req

	handler.MakePayment(c)

	assert.Equal(t, http.StatusOK, w.Code)
}
