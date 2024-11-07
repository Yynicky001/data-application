class TalentRankCalculator:
    def __init__(self, uri, user, password):
        self.driver = GraphDatabase.driver(uri, auth=(user, password))

    def close(self):
        self.driver.close()

    def create_graph(self, developers, projects, contributions):
        with self.driver.session() as session:
            # 清空已有的图
            session.run("MATCH (n) DETACH DELETE n")
            
            # 创建开发者节点
            for dev in developers:
                session.run(
                    """
                    CREATE (:Developer {id: $id, name: $name, followers: $followers})
                    """, id=dev['id'], name=dev['name'], followers=dev['followers_num']
                )

            # 创建项目节点
            for proj in projects:
                session.run(
                    """
                    CREATE (:Project {id: $id, stars: $stars})
                    """, id=proj['id'], stars=proj['stars']
                )

            # 创建贡献关系
            for contrib in contributions:
                session.run(
                    """
                    MATCH (d:Developer {id: $dev_id}), (p:Project {id: $project_id})
                    CREATE (d)-[:CONTRIBUTED {commits: $commits, prs: $prs, issues: $issues}]->(p)
                    """, dev_id=contrib['dev_id'], project_id=contrib['project_id'],
                    commits=contrib['commits'], prs=contrib['prs'], issues=contrib['issues']
                )

    def calculate_page_rank(self):
        with self.driver.session() as session:
            # 使用PageRank算法计算开发者的得分
            result = session.run(
                """
                CALL gds.pageRank.stream({
                    nodeProjection: 'Developer',
                    relationshipProjection: {
                        CONTRIBUTED: {
                            type: 'CONTRIBUTED',
                            properties: 'weight'
                        }
                    },
                    dampingFactor: 0.85
                })
                YIELD nodeId, score
                RETURN gds.util.asNode(nodeId).id AS id, score
                ORDER BY score DESC
                """
            )
            return [{"dev_id": record["id"], "talent_rank": record["score"]} for record in result]