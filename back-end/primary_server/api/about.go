package api

import (
	"github-data-evaluator/service"
	"github.com/gin-gonic/gin"
)

func About() gin.HandlerFunc {
	return func(c *gin.Context) {
		abouts, err := service.GetAboutService(c).GetAbout()
		if err != nil {
			c.JSON(500, gin.H{"error": "get about error"})
			return
		}
		c.JSON(200, gin.H{"data": abouts})
	}
}
