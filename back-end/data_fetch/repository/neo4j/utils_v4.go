package neo4j

import (
	"data_fetch/pkg/utils"
	"data_fetch/repository/model"
)

func GenerateUserCypher() string {
	return `CREATE (u: user {id: $id, login: $login, avatar_url: $avatar_url, followers_num: $followers_num, following_num: $following_num, blog: $blog, company: $company, location: $location, html_url: $html_url, created_at: $created_at})`
}

// CreateUserNode creates a node in the graph database for a given user
func CreateUserNode(user *model.User) error {
	// define cypher query to create the node
	cypher := `CREATE (u: user {id: $id, login: $login, avatar_url: $avatar_url, followers_num: $followers_num, following_num: $following_num, blog: $blog, company: $company, location: $location, html_url: $html_url, created_at: $created_at})`

	// run the query
	_, err := session.Run(cypher, user.User2Map())

	if err != nil {
		utils.GetLogger().Errorf("Error creating node:", err)
		return err
	}
	return nil
}

// CreateRepositoryNode creates a node in the graph database for a given repository
func CreateRepositoryNode(repo *model.Repository) error {
	// define cypher query to create the node
	cypher := `CREATE (r: repository {id: $id, name: $name, description: $description, language: $language, created_at: $created_at, updated_at: $updated_at, watchers: $watchers, forks: $forks, stars: $stars, issues: $issues, html_url: $html_url, owner_id: $owner_id}) `

	// run the query
	_, err := session.Run(cypher, repo.Repo2Map())

	if err != nil {
		utils.GetLogger().Errorf("Error creating node:", err)
		return err
	}
	return nil
}

// CreateOrganizationNode creates a node in the graph database for a given organization
func CreateOrganizationNode(org *model.Organization) error {
	// define cypher query to create the node
	cypher := `CREATE (o: organization {id: $id, login: $login, avatar_url: $avatar_url, description: $description, created_at: $created_at})`
	// run the query
	_, err := session.Run(cypher, org.Organization2Map())
	if err != nil {
		utils.GetLogger().Errorf("Error creating node:", err)
		return err
	}
	return nil
}

func CreateConnection(userID, repoID int64) error {
	cypher := `MATCH (u: user),(r: repository) WHERE u.id=$userID AND r.id=$repoID CREATE (u)-[:CONTRIBUTES_TO]->(r)`
	_, err := session.Run(cypher, map[string]interface{}{"userID": userID, "repoID": repoID})
	if err != nil {
		utils.GetLogger().Errorf("Error creating connection:", err)
		return err
	}
	return nil
}
