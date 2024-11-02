package model

type DeveloperDetail struct {
	ID                    int64   `json:"id"`                             //id
	Login                 string  `json:"login" gorm:"unique"`            //用户名
	Name                  string  `json:"name"`                           //用户别名
	HTMLURL               string  `json:"html_url" gorm:"not null"`       //用户主页
	AvatarURL             string  `json:"avatar_url"`                     //头像url
	FollowersNum          int     `json:"followers_num" gorm:"default:0"` //获取用户粉丝数量
	FollowingNum          int     `json:"following_url" gorm:"default:0"` //获取用户正在关注的用户数量
	Blog                  string  `json:"blog"`                           //博客
	Location              string  `json:"location"`                       //位置
	Company               string  `json:"company"`                        //公司
	Talent                float64 `json:"talent"`                         //代码能力评级
	RankNum               int     `json:"rank_num"`                       //排名
	Nation                string  `json:"nation"`                         //猜测的国籍
	ConfidenceCoefficient float64 `json:"confidence_coefficient"`         //国籍置信度
	Domain                string  `json:"domain"`                         //领域
}
