package model

type Repo struct {
	ID       int64  `json:"id"`       // 仓库id
	Name     string `json:"name"`     // 仓库名
	OwnerID  int64  `json:"owner_id"` // 仓库拥有者id
	Language string `json:"language"` //代码量最多的语言
	Stars    int    `json:"stars"`    // star数量
	Forks    int    `json:"forks"`    // fork数量
	Issue    int    `json:"issues"`   // issue数量
	HTMLURL  string `json:"html_url"` // 仓库地址

	CreatedAt    string         `json:"created_at"` // 创建时间
	UpdatedAt    string         `json:"updated_at"` // 更新时间
	Contribution []Contribution `json:"contribution" `
}
