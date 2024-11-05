package dao

import (
	"context"
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/repository/model"
	"github-data-evaluator/repository/mysql/db"
	"gorm.io/gorm"
	"sync"
)

var (
	developerDaoIns  *DeveloperDao
	developerDaoOnce sync.Once
)

type DeveloperDao struct {
	*gorm.DB
}

func GetDeveloperDao() *DeveloperDao {
	developerDaoOnce.Do(func() {
		developerDaoIns = &DeveloperDao{db.NewDBClient(context.Background())}
	})
	return developerDaoIns
}

func (dd *DeveloperDao) BatchInsert(developers []*model.Developer) (err error) {
	// 执行批量插入，每次插入100条记录
	batchSize := 100
	err = dd.DB.CreateInBatches(developers, batchSize).Error
	if err != nil {
		utils.GetLogger().Errorf("Error inserting developers:", err)
	}
	return
}
