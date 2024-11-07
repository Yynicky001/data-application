package dao

import (
	"context"
	"github-data-evaluator/repository/model"
	"github-data-evaluator/repository/mysql/db"
	"gorm.io/gorm"
	"sync"
)

var aboutDaoIns *AboutDao
var aboutDaoOnce sync.Once

type AboutDao struct {
	db *gorm.DB
}

func GetAboutDao(c context.Context) *AboutDao {
	if aboutDaoIns == nil {
		aboutDaoOnce.Do(func() {
			aboutDaoIns = &AboutDao{db: db.NewDBClient(c)}
		})
	}
	return aboutDaoIns
}

func (a *AboutDao) QueryAbout() (abouts []*model.About, err error) {
	err = a.db.Model(&model.About{}).Find(&abouts).Error
	return
}

func (a *AboutDao) InsertAbouts(about []*model.About) (err error) {
	return a.db.Model(&model.About{}).Create(&about).Error
}
