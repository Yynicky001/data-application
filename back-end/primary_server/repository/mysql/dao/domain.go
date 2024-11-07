package dao

import (
	"context"
	"github-data-evaluator/repository/mysql/db"
	"gorm.io/gorm"
	"sync"
)

type Domain struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

var domainDaoOnce sync.Once
var domainDaoIns *DomainDao

type DomainDao struct {
	*gorm.DB
}

func GetDomainDao() *DomainDao {
	domainDaoOnce.Do(func() {
		domainDaoIns = &DomainDao{
			DB: db.NewDBClient(context.Background()),
		}
	})
	return domainDaoIns
}

func (d *DomainDao) GetDomainList() (domains []*Domain, err error) {
	err = d.Model(&Domain{}).Find(&domains).Error
	return
}
