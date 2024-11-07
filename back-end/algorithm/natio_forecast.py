import langdetect
from polyglot.detect import Detector
from neo4j import GraphDatabase
from flask import Flask, jsonify, request

app = Flask(__name__)

# Neo4j驱动初始化
neo4j_driver = GraphDatabase.driver("bolt://localhost:7687", auth=("neo4j", "password"))

class Developer:
    def __init__(self, id, name, followers, nation=None, organization=None, language_confidence=None):
        self.id = id
        self.name = name
        self.followers = followers
        self.nation = nation
        self.organization = organization
        self.language_confidence = language_confidence  # 语言检测置信度

# 语言检测函数
def enhanced_language_detection(text):
    try:
        detector = Detector(text)
        languages = detector.languages
        # 获取最可能的语言及置信度
        if languages:
            primary_lang = languages[0].code
            confidence = languages[0].confidence
            return primary_lang, confidence
    except:
        return "unknown", 0.0

# Neo4j构建关系网络
def create_complex_relationships_in_neo4j(dev_id, profile_nation, organization, company, email_domain, secondary_connections):
    with neo4j_driver.session() as session:
        session.run("""
            MERGE (d:Developer {id: $id})
            SET d.nation = $nation, d.organization = $organization, d.company = $company, d.email_domain = $email_domain
        """, id=dev_id, nation=profile_nation, organization=organization, company=company, email_domain=email_domain)
        
        # 建立二级连接关系
        for conn in secondary_connections:
            session.run("""
                MERGE (d2:Developer {id: $conn_id})
                MERGE (d1:Developer {id: $dev_id})
                MERGE (d1)-[:CONNECTED_TO]->(d2)
            """, dev_id=dev_id, conn_id=conn["id"])

# 使用加权模型推断开发者Nation
def infer_nation_with_weights(developer, dev_connections):
    confidence_scores = {}

    # 语言置信度
    lang, lang_conf = enhanced_language_detection(developer.language_confidence)
    
    # 根据语言加权分配置信度
    if lang in ["zh-cn", "zh-tw"]:
        confidence_scores["China"] = confidence_scores.get("China", 0) + 0.4 * lang_conf
    elif lang == "en":
        confidence_scores["USA"] = confidence_scores.get("USA", 0) + 0.2 * lang_conf
    elif lang == "es":
        confidence_scores["Spain"] = confidence_scores.get("Spain", 0) + 0.3 * lang_conf
    elif lang == "fr":
        confidence_scores["France"] = confidence_scores.get("France", 0) + 0.3 * lang_conf
    elif lang == "ar":
        confidence_scores["UAE"] = confidence_scores.get("UAE", 0) + 0.4 * lang_conf
    elif lang == "ru":
        confidence_scores["Russia"] = confidence_scores.get("Russia", 0) + 0.3 * lang_conf
    elif lang == "de":
        confidence_scores["Germany"] = confidence_scores.get("Germany", 0) + 0.3 * lang_conf
    elif lang == "pt":
        confidence_scores["Brazil"] = confidence_scores.get("Brazil", 0) + 0.35 * lang_conf
    elif lang == "it":
        confidence_scores["Italy"] = confidence_scores.get("Italy", 0) + 0.3 * lang_conf
    elif lang == "ja":
        confidence_scores["Japan"] = confidence_scores.get("Japan", 0) + 0.4 * lang_conf
    elif lang == "ko":
        confidence_scores["South Korea"] = confidence_scores.get("South Korea", 0) + 0.4 * lang_conf

    # 社交连接加权
    for connection in dev_connections:
        if connection.get("nation"):
            confidence_scores[connection.get("nation")] = confidence_scores.get(connection.get("nation"), 0) + 0.2

    # 公司和邮箱域置信度
    org_country = get_organization_country(developer.organization)
    if org_country:
        confidence_scores[org_country] = confidence_scores.get(org_country, 0) + 0.3

    if confidence_scores:
        return max(confidence_scores, key=confidence_scores.get), max(confidence_scores.values())
    return None, 0.0

# 获取组织的国别信息
def get_organization_country(organization):
    organization_map = {
         
        # 中国 (China)
        "Tencent": "China",
        "Alibaba": "China",
        "Baidu": "China",
        "Huawei": "China",
        "ByteDance": "China",
        "JD.com": "China",
        "Xiaomi": "China",
        "NetEase": "China",
        "Weibo": "China",
        "Didi Chuxing": "China",
        "Ping An": "China",
        "Lenovo": "China",
        "Meituan": "China", 

         # 美国 (USA)
        "Google": "USA",
        "Microsoft": "USA",
        "Apple": "USA",
        "Amazon": "USA",
        "Facebook": "USA",
        "Intel": "USA",
        "IBM": "USA",
        "Oracle": "USA",
        "Twitter": "USA",
        "LinkedIn": "USA",
        "Uber": "USA",
        "Tesla": "USA",
        "SpaceX": "USA",
        "Salesforce": "USA",
        "Slack": "USA",
        "Red Hat": "USA"
    }
    return organization_map.get(organization, None)

# Flask API 路由
@app.route('/process_data', methods=['POST'])
def process_data():
    data = request.get_json()
    developers = data.get("developers", [])
    dev_connections = data.get("connections", [])
    results = []

    for dev_data in developers:
        developer = Developer(
            id=dev_data["id"],
            name=dev_data["name"],
            followers=dev_data["followers_num"],
            nation=dev_data.get("nation"),
            organization=dev_data.get("organization"),
            language_confidence=dev_data.get("readme_text", "")
        )
        
        # Neo4j 中创建关系网络
        create_complex_relationships_in_neo4j(
            dev_id=developer.id,
            profile_nation=developer.nation,
            organization=developer.organization,
            company=dev_data.get("company"),
            email_domain=dev_data.get("email").split('@')[-1] if dev_data.get("email") else None,
            secondary_connections=dev_data.get("secondary_connections", [])
        )

        # Nation 推测与置信度
        inferred_nation, confidence = infer_nation_with_weights(developer, dev_connections)
        results.append({
            "id": developer.id,
            "inferred_nation": inferred_nation,
            "confidence": confidence
        })

    return jsonify(results)

if __name__ == "__main__":
    app.run(port=5000)