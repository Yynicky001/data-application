package dao

import (
	"context"
	"github-data-evaluator/repository/model"
	"github-data-evaluator/repository/mysql/db"
	"gorm.io/gorm"
	"sync"
)

type RankDao struct {
	db *gorm.DB
}

var rankDaoOnce sync.Once
var rankDaoIns *RankDao

func GetRankDao(c context.Context) *RankDao {
	if rankDaoIns == nil {
		rankDaoOnce.Do(func() {
			rankDaoIns = &RankDao{db: db.NewDBClient(c)}
		})
	}
	return rankDaoIns
}

func (r *RankDao) QueryRankPages(pageSize, offset int) (ranks []*model.Rank, err error) {
	err = r.db.Model(&model.Rank{}).Order("rank_num DESC").Limit(pageSize).Offset(offset).Find(&ranks).Error
	return
}
