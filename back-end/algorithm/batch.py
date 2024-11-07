def calculate_page_rank(self):
    with self.driver.session() as session:
       
        # 删除旧图
        session.run("CALL gds.graph.drop('devGraph', false) YIELD graphName")

        # 创建新图
        session.run(
            """
            CALL gds.graph.project(
                'devGraph',
                'Developer',
                {
                    CONTRIBUTED: {
                        type: 'CONTRIBUTED',
                        properties: 'weight'
                    }
                }
            )
            """
        )

        # PageRank
        session.run(
            """
            CALL gds.pageRank.write('devGraph', {
                dampingFactor: 0.85,
                writeProperty: 'talent_rank'
            })
            YIELD nodePropertiesWritten
            """
        )

        # 分批
        talent_ranks = []
        batch_size = 10000
        offset = 0

        while True:
            result = session.run(
                """
                MATCH (d:Developer)
                RETURN d.id AS id, d.talent_rank AS score
                ORDER BY score DESC
                SKIP $offset
                LIMIT $batch_size
                """, offset=offset, batch_size=batch_size
            )

            
            batch = [{"dev_id": record["id"], "talent_rank": record["score"]} for record in result]
            if not batch:  
                break
            talent_ranks.extend(batch)
            offset += batch_size

        # 删除图数据(释放资源)
        session.run("CALL gds.graph.drop('devGraph')")

        return talent_ranks
