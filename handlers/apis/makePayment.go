package apis

import (
	"billingEngine/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (a *ApiHandler) MakePayment(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(400, gin.H{"error": "Missing userId parameter"})
		return
	}
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	amountStr := c.Query("amount")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	loanDataInterface, ok := a.userLoanMap.Load(int(userIdInt))
	if !ok {
		c.JSON(400, gin.H{"error": "Missing userId parameter"})
		return
	}
	loanData := loanDataInterface.(models.Loan)
	amountDurationInWeeks := loanData.LoanEndDate.Sub(loanData.LoanStartDate).Hours() / HoursInAWeek
	totalAmountToBePaid := loanData.AmountTaken + (loanData.AmountTaken * InterestRate / 100)
	weeklyPaymentAmount := totalAmountToBePaid / amountDurationInWeeks
	if amount < weeklyPaymentAmount {
		c.JSON(400, gin.H{"error": fmt.Sprintf("amount should be equal to or more than %.2f", weeklyPaymentAmount)})
		return
	}
	loanData.AmountPaid += amount
	a.userLoanMap.Store(int(userIdInt), loanData)
	c.JSON(200, gin.H{"payment": "success"})
	return
}
