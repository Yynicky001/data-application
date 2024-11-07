package api

import (
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/repository/mysql/dao"
	"github-data-evaluator/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Rank() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, limit, err := utils.GetPageQuery(c.Query("page"), c.Query("limit"))
		if err != nil {
			// 如果转换失败，返回错误
			c.JSON(http.StatusBadRequest, gin.H{"error": "page query parameter must be an integer"})
			return
		}
		ranks, err := service.GetRankService(c).GetPagesRank(page, limit)
		if err != nil {
			c.JSON(500, gin.H{"error": "get rank by pages error"})
			return
		}
		utils.GetLogger().Infof("Rank: %v", ranks)
		// 返回数据
		c.JSON(200, gin.H{"data": ranks})

	}
}

func RankByDomain() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req map[string]string
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.GetLogger().Errorf("Error binding JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		utils.GetLogger().Infof("Request Body: %v", req)

		page, perPage, err := utils.GetPageQuery(req["page"], req["per_page"])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page or per_page query parameters"})
			return
		}
		domain := req["domain"]
		if domain == "" {
			utils.GetLogger().Errorf("Error domain is empty")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid domain"})
			return
		}

		users, err := service.GetRankService(c).RankPagesByDomain(page, perPage, domain)
		if err != nil {
			utils.GetLogger().Errorf("Error get rank by domain and pages: %v", err)
			c.JSON(500, gin.H{"error": "get rank by domain and pages error"})
			return
		}
		c.JSON(200, gin.H{"data": users})

	}
}

// DomainList 获取所有领域
func DomainList() gin.HandlerFunc {
	return func(c *gin.Context) {
		domains, err := dao.GetDomainDao().GetDomainList()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"data": domains})
	}
}
