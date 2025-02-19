package apis

import (
	"billingEngine/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func (a *ApiHandler) IsDelinquent(c *gin.Context) {
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
	loanDataInterface, ok := a.userLoanMap.Load(int(userIdInt))
	if !ok {
		c.JSON(400, gin.H{"error": "Missing userId parameter"})
		return
	}
	loanData := loanDataInterface.(models.Loan)
	amountDurationInWeeks := loanData.LoanEndDate.Sub(loanData.LoanStartDate).Hours() / HoursInAWeek
	totalAmountToBePaid := loanData.AmountTaken + (loanData.AmountTaken * InterestRate / 100)
	weeklyPaymentAmount := totalAmountToBePaid / amountDurationInWeeks
	paidWeeks := int(loanData.AmountPaid / weeklyPaymentAmount)
	weeksPassed := int(time.Now().Sub(loanData.LoanStartDate).Hours() / HoursInAWeek)
	if (weeksPassed - paidWeeks) > 2 {
		c.JSON(200, gin.H{"IsDelinquent": true})
		return
	}
	c.JSON(200, gin.H{"IsDelinquent": false})
	return
}
