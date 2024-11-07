package model

import (
	"github.com/google/go-github/v66/github"
)

type Developer struct {
	ID        int64  `json:"id"`         //id
	Login     string `json:"login"`      //用户名
	HTMLURL   string `json:"html_url"`   //用户主页
	AvatarURL string `json:"avatar_url"` //头像url
}

func User2Developer(user *github.User) *Developer {
	return &Developer{
		ID:        user.GetID(),
		Login:     user.GetLogin(),
		HTMLURL:   user.GetHTMLURL(),
		AvatarURL: user.GetAvatarURL(),
	}
}

func Users2Developers(users []*github.User) (developers []*Developer) {
	for _, user := range users {
		developers = append(developers, User2Developer(user))
	}
	return
}

//// BeforeCreate GORM提供的钩子函数 BeforeCreate 在创建记录之前调用
//func (d *Developer) BeforeCreate(_ *gorm.DB) (err error) {
//	if d.Name == "" {
//		d.Name = d.Login
//	}
//	return
//}
