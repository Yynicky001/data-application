package service

import (
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/repository/dao"
	"github-data-evaluator/repository/model"
	"sync"
)

var developerServiceIns *DeveloperService
var developerServiceOnce sync.Once

type DeveloperService struct {
	devDao *dao.DeveloperDao
}

func GetDeveloperService() *DeveloperService {
	developerServiceOnce.Do(func() {
		developerServiceIns = &DeveloperService{devDao: dao.GetDeveloperDao()}
	})
	return developerServiceIns
}

// BatchInsertDevelopers 批量存储开发者数据
func (d *DeveloperService) BatchInsertDevelopers(batch []*model.Developer) {
	utils.GetLogger().Infof("Inserted %d developers into the database.", len(batch))
	if err := d.devDao.BatchInsert(batch); err != nil {
		utils.GetLogger().Errorf("Error inserting developers:", err)
	}
}
