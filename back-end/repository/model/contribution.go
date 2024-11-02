package model

import (
	"github.com/google/go-github/v66/github"
)

type Contribution struct {
	ID           int64 `json:"id" gorm:"primaryKey;autoIncrement"`
	DeveloperID  int64 `json:"developer_id"`
	RepositoryID int64 `json:"repository_id"`
	Contribution int   `json:"contribution"`
}

func Conversion2Contribution(contribution *github.Contributor, repoId int64) *Contribution {
	return &Contribution{
		DeveloperID:  contribution.GetID(),
		RepositoryID: repoId,
		Contribution: contribution.GetContributions(),
	}
}
