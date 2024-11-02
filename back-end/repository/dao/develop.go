package dao

import (
	"context"
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/repository/db"
	"github-data-evaluator/repository/model"
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
	// 开启事务
	tx := dd.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// 执行批量插入，每次插入100条记录
	batchSize := 100
	if err := tx.CreateInBatches(developers, batchSize).Error; err != nil {
		// 如果有错误发生，回滚事务
		utils.LogrusObj.Errorln("Error inserting developers:", err)
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}
