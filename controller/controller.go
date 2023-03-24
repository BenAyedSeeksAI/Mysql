package controller

import (
	"mysqlCountryProjects/apires"
	"mysqlCountryProjects/db"

	"github.com/gin-gonic/gin"
)

func GetGovernmentsPerCountryStats(c *gin.Context) {
	data, err := db.GetGovernmentsPerCountryStats()
	if err != nil {
		apires.DefaultError(c)
		return
	}
	apires.Success(c, 200, data)
}
func GetCountryPerContinentStats(c *gin.Context) {
	data, err := db.GetCountryPerContinentStats()
	if err != nil {
		apires.DefaultError(c)
		return
	}
	apires.Success(c, 200, data)
}
