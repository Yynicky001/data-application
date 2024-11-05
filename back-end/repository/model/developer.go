package model

import (
	"github.com/google/go-github/v66/github"
	"gorm.io/gorm"
)

type Developer struct {
	ID           int64          `json:"id" gorm:"primaryKey"` //id
	DeveloperID  int64          `json:"developer_id" gorm:"unique"`
	Login        string         `json:"login"`                          //用户名
	Name         string         `json:"name"`                           //用户别名
	HTMLURL      string         `json:"html_url" gorm:"not null"`       //用户主页
	AvatarURL    string         `json:"avatar_url"`                     //头像url
	FollowersNum int            `json:"followers_num" gorm:"default:0"` //获取用户粉丝数量
	FollowingNum int            `json:"following_url" gorm:"default:0"` //获取用户正在关注的用户数量
	Blog         string         `json:"blog"`                           //博客
	Location     string         `json:"location"`                       //位置
	Company      string         `json:"company"`                        //公司
	CreatAt      string         `json:"created_at"`                     //创建时间
	Contribution []Contribution `json:"contribution" gorm:"foreignKey:DeveloperID"`
}

func User2Developer(user *github.User) *Developer {
	return &Developer{
		DeveloperID:  user.GetID(),
		Login:        user.GetLogin(),
		Name:         user.GetName(),
		HTMLURL:      user.GetHTMLURL(),
		AvatarURL:    user.GetAvatarURL(),
		FollowersNum: user.GetFollowers(),
		FollowingNum: user.GetFollowing(),
		Blog:         user.GetBlog(),
		Location:     user.GetLocation(),
		Company:      user.GetCompany(),
		CreatAt:      user.GetCreatedAt().String(),
	}
}

// BeforeCreate GORM提供的钩子函数 BeforeCreate 在创建记录之前调用
func (d *Developer) BeforeCreate(tx *gorm.DB) (err error) {
	if d.Name == "" {
		d.Name = d.Login
	}
	return
}
