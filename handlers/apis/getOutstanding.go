package apis

import (
	"billingEngine/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (a *ApiHandler) GetOutStandingAmount(c *gin.Context) {
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
	totalAmountToBePaid := loanData.AmountTaken + (loanData.AmountTaken * InterestRate / 100)
	outStandingAmount := totalAmountToBePaid - loanData.AmountPaid
	c.JSON(200, gin.H{"outStandingAmount": outStandingAmount})
	return
}
