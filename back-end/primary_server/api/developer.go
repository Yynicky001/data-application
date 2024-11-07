package api

import (
	"context"
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/repository/model"
	"github-data-evaluator/service"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v66/github"
	"net/http"
)

func Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		search := c.Query("query")
		page, perPage, err := utils.GetPageQuery(c.Query("page"), c.Query("per_page"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page or per_page query parameters"})
			return
		}
		opts := &github.SearchOptions{
			ListOptions: github.ListOptions{
				PerPage: perPage,
				Page:    page,
			},
		}
		searchRes, _, err := github.NewClient(&http.Client{}).Search.Users(context.Background(), search, opts)
		if err != nil {
			utils.GetLogger().Errorf("Error fetching users: %v", err)
			c.JSON(500, gin.H{
				"error": "Error search users",
			})
		}
		developers := model.Users2Developers(searchRes.Users)

		c.JSON(200, gin.H{
			"data": gin.H{
				"users": developers,
				"total": searchRes.GetTotal(),
			},
		})
	}
}

func DeveloperDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		login := c.Query("login")
		if login == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "login query parameter is required"})
			return
		}
		developer, err := service.GetDeveloperDetailService(c).GetDeveloperDetail(login)
		if err != nil {
			c.JSON(500, gin.H{"error": "get developer detail error"})
			return
		}
		c.JSON(200, gin.H{"data": developer})
	}
}
