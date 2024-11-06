package router

import (
	"github-data-evaluator/api"
	"github-data-evaluator/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	limiter := middleware.NewIPRateLimiter()
	ginRouter.Use(middleware.Cors(), middleware.RateLimitByIP(limiter))
	ginRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	v1 := ginRouter.Group("/v1/github")

	v1.GET("/rank", api.Rank())

	v1.GET("/domain/rank", api.RankByDomain())

	v1.GET("/developer/details", api.DeveloperDetail())

	v1.POST("/search", api.Search())

	v1.POST("/about_us", api.About())

	return ginRouter
}
