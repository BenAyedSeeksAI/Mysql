package main

import (
	"mysqlCountryProjects/controller"

	"github.com/gin-gonic/gin"
)

func routes() {
	r := gin.Default()
	r.GET("/getGov", controller.GetGovernmentsPerCountryStats)
	r.GET("/getContinent", controller.GetCountryPerContinentStats)
	r.Run(":8081")
}

func main() {
	routes()
}
 