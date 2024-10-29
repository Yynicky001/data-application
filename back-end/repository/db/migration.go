package db

import "github-data-evaluator/repository/model"

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.Developer{}, &model.DeveloperRank{})
	if err != nil {
		panic(err)
	}
}
