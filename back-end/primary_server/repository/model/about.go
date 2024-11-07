package model

type About struct {
	Login     string `json:"login"`      //用户名
	HTMLURL   string `json:"htmlurl"`    //网址
	AvatarURL string `json:"avatar_url"` //头像url

}
