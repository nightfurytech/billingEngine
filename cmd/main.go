package main

import (
	"billingEngine/handlers/apis"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"time"

	"billingEngine/models"
)

var (
	now              = time.Now()
	fiftyWeeksBefore = now.AddDate(0, 0, -350)
	tenWeeksBefore   = now.AddDate(0, 0, -70)
	fortyWeeksAfter  = now.AddDate(0, 0, 280)
	loans            = []models.Loan{
		{UserID: 1, LoanStartDate: fiftyWeeksBefore, LoanEndDate: now, AmountTaken: 5000000, AmountPaid: 50 * 110000},
		{UserID: 2, LoanStartDate: fiftyWeeksBefore, LoanEndDate: now, AmountTaken: 5000000, AmountPaid: 45 * 110000},
		{UserID: 3, LoanStartDate: fiftyWeeksBefore, LoanEndDate: now, AmountTaken: 5000000, AmountPaid: 0},
		{UserID: 4, LoanStartDate: fiftyWeeksBefore, LoanEndDate: now, AmountTaken: 5000000, AmountPaid: 49 * 110000},
		{UserID: 5, LoanStartDate: tenWeeksBefore, LoanEndDate: fortyWeeksAfter, AmountTaken: 5000000, AmountPaid: 7 * 110000},
	}
)

func main() {
	logger, err := zap.NewProduction() // Use zap.NewDevelopment() for development mode
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	apiHandler := apis.NewApiHandler(logger, loans)
	r := gin.Default()
	r.GET("/getOutstandingAmount/:userId", apiHandler.GetOutStandingAmount)
	r.GET("/delinquent/:userId", apiHandler.IsDelinquent)
	r.POST("/payment/:userId", apiHandler.MakePayment)
	log.Println("Server is serving at port 8080")
	r.Run(":8080")
}
