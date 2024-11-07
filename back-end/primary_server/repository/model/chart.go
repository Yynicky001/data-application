package model

type Chart struct {
	Id              int              `json:"id"`
	Login           string           `json:"login"`
	Date            string           `json:"date"`
	Contributions   string           `json:"contributions"`
	RepoLanguageTop []*ChartLanguage `json:"repoLanguageTop"`
	CommitThisYear  int              `json:"commit_this_year"`
	TotalStars      int              `json:"totalStars"`
	PRs             int              `json:"PRs"`
	Issues          int              `json:"issues"`
	ContributionTo  int              `json:"contributionTo"`
	CommitDate      string           `json:"commitDate"`
	Commits         string           `json:"commits"`
}

type ChartData struct {
	Id             int    `json:"id"`
	Login          string `json:"login"`
	Date           string `json:"date"`
	Contributions  string `json:"contributions"`
	CommitThisYear int    `json:"commit_this_year"`
	TotalStars     int    `json:"totalStars"`
	PRs            int    `json:"PRs"`
	Issues         int    `json:"issues"`
	ContributionTo int    `json:"contributionTo"`
	CommitDate     string `json:"commitDate"`
	Commits        string `json:"commits"`
}

type ChartLanguage struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Language string `json:"language"`
	Value    int    `json:"value"`
}
