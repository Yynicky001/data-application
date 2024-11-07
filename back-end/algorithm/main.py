def main():
    # 从MySQL读取数据
    developers = get_mysql_data("SELECT * FROM developers")
    projects = get_mysql_data("SELECT * FROM projects")
    contributions = get_mysql_data("SELECT * FROM contributions")

    # 创建Neo4j图
    rank_calculator = TalentRankCalculator(neo4j_uri, neo4j_user, neo4j_password)
    rank_calculator.create_graph(developers, projects, contributions)
    
    # 计算PageRank
    rank_results = rank_calculator.calculate_page_rank()
    rank_calculator.close()

    # 保存结果到MySQL
    conn = mysql.connector.connect(**mysql_config)
    cursor = conn.cursor()
    
    # 更新开发者排名
    cursor.execute("DELETE FROM developer_rank")  # 清除旧数据
    for result in rank_results:
        cursor.execute(
            "INSERT INTO developer_rank (dev_id, talent_rank) VALUES (%s, %s)",
            (result['dev_id'], result['talent_rank'])
        )
    
    conn.commit()
    cursor.close()
    conn.close()
    print("Developer TalentRank scores have been updated in MySQL.")

if __name__ == "__main__":
    main()