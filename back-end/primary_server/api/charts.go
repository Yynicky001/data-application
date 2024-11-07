package api

import (
	"github-data-evaluator/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Charts() gin.HandlerFunc {
	return func(c *gin.Context) {
		login := c.Query("login")
		if login == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "login query parameter is required"})
			return
		}
		data, err := service.GetChartService(c).GetChartData(login)
		if err != nil {
			c.JSON(500, gin.H{"error": "get chart data error"})
			return
		}
		c.JSON(200, gin.H{"data": data})
	}
}
