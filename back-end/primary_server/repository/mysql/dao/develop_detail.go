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

func (d *DeveloperDetailDao) QueryDeveloperDetail(id int64) (developerDetail *model.DeveloperDetail, err error) {
	err = d.db.Where("id = ?", id).First(&developerDetail).Error
	return
}
