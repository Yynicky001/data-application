package dao

import (
	"github-data-evaluator/repository/model"
	"github-data-evaluator/repository/mysql/db"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"sync"
)

var chartDaoIns *ChartDao
var chartDaoOnce sync.Once

type ChartDao struct {
	db *gorm.DB
}

func GetChartDao(c context.Context) *ChartDao {
	if chartDaoIns == nil {
		chartDaoOnce.Do(func() {
			chartDaoIns = &ChartDao{db: db.NewDBClient(c)}
		})
	}
	return chartDaoIns
}

func (c *ChartDao) QueryChartData(login string) (charts *model.ChartData, err error) {
	err = c.db.Model(&model.ChartData{}).Where("login = ?", login).First(&charts).Error
	return
}

func (c *ChartDao) QueryChartLanguages(login string) (charts []*model.ChartLanguage, err error) {
	err = c.db.Model(&model.ChartLanguage{}).Where("login = ?", login).Find(&charts).Error
	return
}
