package service

import (
	"context"
	"github-data-evaluator/repository/model"
	"github-data-evaluator/repository/mysql/dao"
	"sync"
)

var developerDetailServiceIns *DeveloperDetailService
var developerDetailServiceOnce sync.Once

type DeveloperDetailService struct {
	Dao *dao.DeveloperDetailDao
}

func GetDeveloperDetailService(c context.Context) *DeveloperDetailService {
	developerDetailServiceOnce.Do(func() {
		developerDetailServiceIns = &DeveloperDetailService{
			Dao: dao.GetDeveloperDetailDao(c),
		}
	})
	return developerDetailServiceIns
}

func (d *DeveloperDetailService) GetDeveloperDetail(login string) (developerDetail *model.DeveloperDetail, err error) {
	return d.Dao.QueryDeveloperDetail(login)
}
