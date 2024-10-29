package model

import "github.com/google/go-github/v33/github"

type Developer struct {
	ID           int64  `json:"id"`            //id
	Login        string `json:"login"`         //用户名
	Name         string `json:"name"`          //用户别名
	HTMLURL      string `json:"html_url"`      //用户主页
	AvatarURL    string `json:"avatar_url"`    //头像url
	FollowersNum int    `json:"followers_num"` //获取用户粉丝数量
	FollowingNum int    `json:"following_url"` //获取用户正在关注的用户数量
	Blog         string `json:"blog"`          //博客
	Location     string `json:"location"`      //位置
	Company      string `json:"company"`       //公司
}

func (d *Developer) User2Developer(user *github.User) *Developer {
	d.ID = user.GetID()
	d.Login = user.GetLogin()
	d.Name = user.GetName()
	d.HTMLURL = user.GetHTMLURL()
	d.AvatarURL = user.GetAvatarURL()
	d.FollowersNum = user.GetFollowers()
	d.FollowingNum = user.GetFollowing()
	d.Blog = user.GetBlog()
	d.Location = user.GetLocation()
	d.Company = user.GetCompany()
	return d
}
