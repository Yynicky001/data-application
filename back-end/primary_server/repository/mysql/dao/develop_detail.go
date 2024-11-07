package dao

import (
	"context"
	"github-data-evaluator/repository/model"
	"github-data-evaluator/repository/mysql/db"
	"gorm.io/gorm"
	"sync"
)

var developerDetailDaoIns *DeveloperDetailDao
var developerDetailDaoOnce sync.Once

type DeveloperDetailDao struct {
	db *gorm.DB
}

func GetDeveloperDetailDao(c context.Context) *DeveloperDetailDao {
	developerDetailDaoOnce.Do(func() {
		developerDetailDaoIns = &DeveloperDetailDao{
			db: db.NewDBClient(c),
		}
	})
	return developerDetailDaoIns
}

func (d *DeveloperDetailDao) QueryDeveloperDetail(login string) (developerDetail *model.DeveloperDetail, err error) {
	err = d.db.Where("login = ?", login).First(&developerDetail).Error
	return
}
