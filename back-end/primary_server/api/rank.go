package api

import (
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Rank() gin.HandlerFunc {
	return func(c *gin.Context) {
		pageParam := c.Query("page")
		// 默认第一页
		if pageParam == "" {
			pageParam = "1"
		}
		limitParam := c.Query("limit")
		// 默认每页10条数据
		if limitParam == "" {
			limitParam = "10"
		}
		// 将字符串转换为int
		page, err := strconv.Atoi(pageParam)
		if err != nil {
			// 如果转换失败，返回错误
			c.JSON(http.StatusBadRequest, gin.H{"error": "page query parameter must be an integer"})
			return
		}
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			// 如果转换失败，返回错误
			c.JSON(http.StatusBadRequest, gin.H{"error": "page query parameter must be an integer"})
			return
		}
		ranks, err := service.GetRankService(c).GetRankBy(page, limit)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		utils.GetLogger().Infof("Rank: %v", ranks)
		// 返回数据
		c.JSON(200, gin.H{"data": ranks})

	}
}

func RankByDomain() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
