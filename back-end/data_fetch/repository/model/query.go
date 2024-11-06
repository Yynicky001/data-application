package model

import "github.com/shurcooL/githubv4"

type UserQuery struct {
	Search struct {
		Nodes    []User
		PageInfo struct {
			EndCursor   githubv4.String
			HasNextPage githubv4.Boolean
		}
	} `graphql:"search(type: USER, query: $query, first: $first, after: $cursor)"`
}

type User struct {
	ID                        githubv4.ID
	Login                     string
	AvatarURL                 string
	Blog                      string
	Company                   string
	Location                  string
	URL                       string
	CreatedAt                 string
	Followers                 struct{ TotalCount int }
	Following                 struct{ TotalCount int }
	Organizations             Organizations
	RepositoriesContributedTo Repositories
}

func (u *User) User2Map() map[string]interface{} {
	return map[string]interface{}{
		"id":            u.ID,
		"login":         u.Login,
		"avatar_url":    u.AvatarURL,
		"followers_num": u.Followers.TotalCount,
		"following_num": u.Following.TotalCount,
		"blog":          u.Blog,
		"company":       u.Company,
		"location":      u.Location,
		"html_url":      u.URL,
		"created_at":    u.CreatedAt,
	}

}

type Organizations struct {
	TotalCount int
	Edges      []struct {
		Node Organization
	}
}

type Organization struct {
	ID          githubv4.ID
	Login       string
	AvatarURL   string
	Description string
	Location    string
	CreatedAt   string
}

func (o *Organization) Organization2Map() map[string]interface{} {
	return map[string]interface{}{
		"id":          o.ID,
		"login":       o.Login,
		"avatar_url":  o.AvatarURL,
		"location":    o.Location,
		"description": o.Description,
		"created_at":  o.CreatedAt,
	}
}

type Repositories struct {
	TotalCount int
	Edges      []struct {
		Node Repository
	}
}

type Repository struct {
	ID    githubv4.ID
	Name  string
	Owner struct {
		ID    githubv4.ID
		login string
	}
	homepageUrl string
	CreatedAt   string
	UpdatedAt   string
	Description string
	Stargazers  struct {
		TotalCount int
	}
	Forks struct {
		TotalCount int
	}
	Watchers struct {
		TotalCount int
	}
	Languages struct {
		Edges []struct {
			Node struct {
				Name string
			}
		}
	}
}

func (r *Repository) Repo2Map() map[string]interface{} {
	var languages []string
	for i, v := range r.Languages.Edges {
		languages[i] = v.Node.Name
	}
	return map[string]interface{}{
		"id":          r.ID,
		"name":        r.Name,
		"owner_id":    r.Owner.ID,
		"stars":       r.Stargazers.TotalCount,
		"forks":       r.Forks.TotalCount,
		"watchers":    r.Watchers.TotalCount,
		"language":    languages,
		"html_url":    r.homepageUrl,
		"created_at":  r.CreatedAt,
		"updated_at":  r.UpdatedAt,
		"description": r.Description,
	}
}
