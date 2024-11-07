package service

import (
	"context"
	"github-data-evaluator/pkg/utils"
	"github-data-evaluator/repository/model"
	"github-data-evaluator/repository/mysql/dao"
	"github.com/google/go-github/v66/github"
	"net/http"
	"sync"
)

var aboutServiceIns *AboutService
var aboutServiceOnce sync.Once

type AboutService struct {
	aboutDao *dao.AboutDao
}

func GetAboutService(c context.Context) *AboutService {
	if aboutServiceIns == nil {
		aboutServiceOnce.Do(func() {
			aboutServiceIns = &AboutService{aboutDao: dao.GetAboutDao(c)}
		})
	}
	return aboutServiceIns
}

func (a *AboutService) GetAbout() (abouts []*model.About, err error) {
	abouts, err = a.aboutDao.QueryAbout()
	if err != nil {
		return nil, err
	}
	if abouts == nil || len(abouts) != 3 {
		abouts = fetchAbouts()
		err := a.aboutDao.InsertAbouts(abouts)
		utils.GetLogger().Errorf("fetch abouts failed, err: %v", err)
	}
	return abouts, nil
}

func fetchAbouts() (abouts []*model.About) {
	client := github.NewClient(&http.Client{})
	get1, _, _ := client.Users.Get(context.Background(), "Yynicky001")
	get2, _, _ := client.Users.Get(context.Background(), "Vegetable-center")
	get3, _, _ := client.Users.Get(context.Background(), "0125nia")

	developers := model.Users2Developers([]*github.User{get1, get2, get3})

	for _, developer := range developers {
		abouts = append(abouts, &model.About{
			AvatarURL: developer.AvatarURL,
			HTMLURL:   developer.HTMLURL,
			Login:     developer.Login,
		})
	}
	return abouts
}
