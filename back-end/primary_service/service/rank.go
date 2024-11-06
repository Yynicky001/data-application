package service

import (
	"context"
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/repository/model"
	"github-data-evaluator/repository/mysql/dao"
	"sync"
)

var rankServiceOnce sync.Once
var rankServiceIns *RankService

type RankService struct {
	*dao.RankDao
}

func GetRankService(c context.Context) *RankService {
	if rankServiceIns == nil {
		rankServiceOnce.Do(func() {
			rankServiceIns = &RankService{dao.GetRankDao(c)}
		})
	}
	return rankServiceIns
}

// GetRankBy 分页获取排行榜
func (r *RankService) GetRankBy(page, limit int) (ranks []*model.Rank, err error) {
	// 计算偏移量
	offset := (page - 1) * limit
	ranks, err = r.RankDao.QueryRankPages(limit, offset)
	if err != nil {
		utils.GetLogger().Errorf("GetRankBy failed, err: %v", err)
		return nil, err
	}
	return ranks, nil
}
