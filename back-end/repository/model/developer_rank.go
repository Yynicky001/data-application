package model

// DeveloperRank 开发者排名结果
type DeveloperRank struct {
	ID                    int64   `json:"id"`                     // 用户ID
	Login                 string  `json:"login"`                  //用户名
	HTMLURL               string  `json:"html_url"`               //用户主页
	AvatarURL             string  `json:"avatar_url"`             //头像url
	Talent                float64 `json:"talent"`                 //代码能力评级
	RankNum               int     `json:"rank_num"`               //排名
	Nation                string  `json:"nation"`                 //猜测的国籍
	ConfidenceCoefficient float64 `json:"confidence_coefficient"` //国籍置信度
	Domain                string  `json:"domain"`                 //领域
}
