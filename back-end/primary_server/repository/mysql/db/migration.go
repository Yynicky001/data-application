package db

import (
	"github-data-evaluator/repository/model"
)

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.Rank{}, &model.DeveloperDetail{}, &model.About{}, &model.ChartData{}, &model.ChartLanguage{})
	if err != nil {
		panic(err)
	}
}
