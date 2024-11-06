package neo4j

import (
	"github-data-evaluator/pkg/utils"
	"github.com/google/go-github/v66/github"
)

// BatchCreateUserNodes creates nodes in the graph database for a batch of users
func BatchCreateUserNodes(usersMap []map[string]interface{}) error {
	// define cypher query to create the nodes
	cypher := `UNWIND $batch AS user
				CREATE (u: user {id: user.id, login: user.login, avatar_url: user.avatar_url, followers_num: user.followers_num, following_num: user.following_num, blog: user.blog, company: user.company, location: user.location, html_url: user.html_url, created_at: user.created_at})`

	// run the query
	_, err := session.Run(cypher, map[string]interface{}{"batch": usersMap})

	if err != nil {
		utils.GetLogger().Errorf("Error creating nodes:", err)
		return err
	}
	return nil
}

func User2Map(user *github.User) map[string]interface{} {
	return map[string]interface{}{
		"id":            user.GetID(),
		"login":         user.GetLogin(),
		"avatar_url":    user.GetAvatarURL(),
		"followers_num": user.GetFollowers(),
		"following_num": user.GetFollowing(),
		"blog":          user.GetBlog(),
		"company":       user.GetCompany(),
		"location":      user.GetLocation(),
		"html_url":      user.GetHTMLURL(),
		"created_at":    user.GetCreatedAt().String(),
		"email":         user.GetEmail(),
	}
}
