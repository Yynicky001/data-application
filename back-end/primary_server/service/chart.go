package service

import (
	"context"
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/repository/model"
	"github-data-evaluator/repository/mysql/dao"
	"sync"
)

var chartServiceIns *ChartService
var ChartServiceOnce sync.Once

type ChartService struct {
	ChartDao *dao.ChartDao
}

func GetChartService(c context.Context) *ChartService {
	if chartServiceIns == nil {
		ChartServiceOnce.Do(func() {
			chartServiceIns = &ChartService{ChartDao: dao.GetChartDao(c)}
		})
	}
	return chartServiceIns
}

func (c *ChartService) GetChartData(login string) (*model.Chart, error) {
	chartData, err := c.ChartDao.QueryChartData(login)
	if err != nil {
		utils.GetLogger().Error("get chart data error")
		return nil, err
	}
	chartLanguages, err := c.ChartDao.QueryChartLanguages(login)
	if err != nil {
		utils.GetLogger().Error("get chart languages error")
		return nil, err
	}

	return &model.Chart{
		Id:              chartData.Id,
		Login:           chartData.Login,
		Commits:         chartData.Commits,
		Contributions:   chartData.Contributions,
		CommitDate:      chartData.CommitDate,
		CommitThisYear:  chartData.CommitThisYear,
		Date:            chartData.Date,
		Issues:          chartData.Issues,
		PRs:             chartData.PRs,
		TotalStars:      chartData.TotalStars,
		ContributionTo:  chartData.ContributionTo,
		RepoLanguageTop: chartLanguages,
	}, nil

}
