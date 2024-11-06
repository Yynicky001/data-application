package model

// Rank 开发者排名结果
type Rank struct {
	ID                    int64   `json:"id"`                         //用户ID
	Login                 string  `json:"login"`                      //用户名
	Talent                float64 `json:"talent"`                     //代码能力评级
	RankNum               int     `json:"rank_num"`                   //排名
	AvatarURL             string  `json:"avatar_url" gorm:"not null"` //头像url
	HTMLURL               string  `json:"html_url" gorm:"not null"`   //用户主页
	Nation                string  `json:"nation"`                     //猜测的国籍
	ConfidenceCoefficient float64 `json:"confidence_coefficient"`     //国籍置信度
	Domain                string  `json:"domain"`                     //领域
}
