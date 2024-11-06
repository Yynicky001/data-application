from neo4j import GraphDatabase
from flask import Flask, jsonify, request
import logging

# 初始化日志
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = Flask(__name__)

# Neo4j驱动初始化
neo4j_driver = GraphDatabase.driver("bolt://localhost:7687", auth=("neo4j", "password"))

class Developer:
    def __init__(self, id, name, followers, nation=None, organization=None, field=None, talent_rank=0):
        self.id = id
        self.name = name
        self.followers = followers
        self.nation = nation
        self.organization = organization
        self.field = field
        self.talent_rank = talent_rank  # TalentRank 分数

# Neo4j 构建多层次关系网络
def create_multilevel_relationships(dev_id, profile_nation, organization, company, email_domain, secondary_connections, level=3):
    """
    建立多层次关系网络，支持到指定层数的连接。
    """
    with neo4j_driver.session() as session:
        # 创建开发者节点并设置基础属性
        session.run("""
            MERGE (d:Developer {id: $id})
            SET d.nation = $nation, d.organization = $organization, d.company = $company, d.email_domain = $email_domain
        """, id=dev_id, nation=profile_nation, organization=organization, company=company, email_domain=email_domain)

        # 递归建立多级关系
        def add_connections(dev_id, connections, depth):
            if depth == 0:
                return
            for conn in connections:
                session.run("""
                    MERGE (d1:Developer {id: $dev_id})
                    MERGE (d2:Developer {id: $conn_id})
                    MERGE (d1)-[:CONNECTED_TO {level: $level}]->(d2)
                """, dev_id=dev_id, conn_id=conn["id"], level=depth)
                
                # 递归添加下一层连接
                deeper_connections = get_developer_connections(conn["id"])
                add_connections(conn["id"], deeper_connections, depth - 1)
        
        add_connections(dev_id, secondary_connections, level)

def get_developer_connections(dev_id):
    """
    获取开发者的直接连接，用于递归构建多级关系。
    假设从数据库查询直接连接的开发者ID列表。
    """
    with neo4j_driver.session() as session:
        results = session.run("""
            MATCH (d:Developer {id: $dev_id})-[:CONNECTED_TO]->(connected:Developer)
            RETURN connected.id AS id
        """, dev_id=dev_id)
        return [{"id": record["id"]} for record in results]
        # 🌰🌰🌰，创建开发者间的多层次关系网络
        # developer_id = "dev123"
        # profile_nation = "China"
        # organization = "TechOrg"
        # company = "TechCorp"
        # mail_domain = "techcorp.com"
        # secondary_connections = [{"id": "dev234"}, {"id": "dev345"}]  初始直接连接
        # create_multilevel_relationships(developer_id, profile_nation, organization, company, email_domain, secondary_connections, level=3)


# 搜索开发者的功能，支持领域、国别和 TalentRank 排序
@app.route('/search_developers', methods=['GET'])
def search_developers():
    field = request.args.get("field")  # 获取查询的领域
    nation = request.args.get("nation")  # 获取可选的国别筛选
    active_since = request.args.get("active_since")  # 获取活跃时间过滤

    # 输入验证
    if not field:
        return jsonify({"error": "Field is required."}), 400

    # Cypher 查询构建
    query = """
        MATCH (d:Developer)
        WHERE d.field = $field
    """

    #  国别筛选
    if nation:
        query += " AND d.nation = $nation"

    # 按 TalentRank 排序，并筛选活跃的开发者
    if active_since:
        query += " AND d.last_active >= $active_since"

    query += " RETURN d ORDER BY d.talent_rank DESC"

    try:
        with neo4j_driver.session() as session:
            results = session.run(query, field=field, nation=nation, active_since=active_since)

            developers = []
            for record in results:
                dev = record["d"]
                developers.append({
                    "id": dev["id"],
                    "name": dev["name"],
                    "followers": dev["followers"],
                    "nation": dev["nation"],
                    "organization": dev["organization"],
                    "field": dev["field"],
                    "talent_rank": dev["talent_rank"]
                })
        return jsonify(developers)  # 返回符合条件的开发者列表
    except Exception as e:
        logger.error(f"查询开发者失败: {e}")
        return jsonify({"error": "An error occurred while querying developers."}), 500

# 更新 TalentRank 分数
def update_talent_rank(developer_id, factors):
    """
    根据一系列因子（followers, contributions, connections 等）更新开发者的 TalentRank。
    """
    base_rank = factors.get("followers", 0) * 0.3 + factors.get("contributions", 0) * 0.5
    network_influence = calculate_network_influence(developer_id)
    talent_rank = base_rank + network_influence

    with neo4j_driver.session() as session:
        session.run("MATCH (d:Developer {id: $id}) SET d.talent_rank = $talent_rank", id=developer_id, talent_rank=talent_rank)

def calculate_network_influence(dev_id):
    """
    计算开发者的网络影响力，可以递归计算多级关系带来的影响。
    """
    with neo4j_driver.session() as session:
        result = session.run("""
            MATCH (d:Developer {id: $id})-[:CONNECTED_TO*..3]-(conn:Developer)
            RETURN SUM(conn.talent_rank) AS influence
        """, id=dev_id)
        influence = result.single()["influence"]
        return influence if influence else 0

# Flask API 路由 - 处理开发者数据
@app.route('/process_data', methods=['POST'])
def process_data():
    data = request.get_json()
    developers = data.get("developers", [])
    results = []

    for dev_data in developers:
        developer = Developer(
            id=dev_data["id"],
            name=dev_data["name"],
            followers=dev_data["followers_num"],
            nation=dev_data.get("nation"),
            organization=dev_data.get("organization"),
            field=dev_data.get("field"),  # 开发者领域
            talent_rank=dev_data.get("talent_rank", 0)  # TalentRank 分数
        )

        # 在 Neo4j 中创建复杂关系网络
        create_multilevel_relationships(
            dev_id=developer.id,
            profile_nation=developer.nation,
            organization=developer.organization,
            company=dev_data.get("company"),
            email_domain=dev_data.get("email").split('@')[-1] if dev_data.get("email") else None,
            secondary_connections=dev_data.get("secondary_connections", []),
            level=3
        )

        # 更新 TalentRank 分数
        factors = {
            "followers": developer.followers,
            "contributions": dev_data.get("contributions", 0)
        }
        update_talent_rank(developer.id, factors)

        results.append({
            "id": developer.id,
            "name": developer.name,
            "nation": developer.nation,
            "organization": developer.organization,
            "field": developer.field,
            "talent_rank": developer.talent_rank
        })

    return jsonify(results)

if __name__ == "__main__":
    app.run(port=5000)