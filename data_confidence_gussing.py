from neo4j import GraphDatabase
from flask import Flask, jsonify, request
import logging

# 初始化日志
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = Flask(__name__)

# Neo4j 驱动初始化
neo4j_driver = GraphDatabase.driver("bolt://localhost:7687", auth=("neo4j", "password"))

class Developer:
    def __init__(self, id, name, followers, nation=None, organization=None, field=None, talent_rank=None, confidence_scores=None):
        self.id = id
        self.name = name
        self.followers = followers
        self.nation = nation
        self.organization = organization
        self.field = field
        self.talent_rank = talent_rank
        self.confidence_scores = confidence_scores or {} # 置信度字典

    def with_confidence_levels(self, confidence_levels):

        """
        根据指定的置信度等级，将数据字段设置为适当的置信度值或 'N/A'。
        """

        return {
            "id": self.id,
            "name": self.name,
            "followers": self.followers,
            "nation": self._mark_confidence_level("nation", confidence_levels),
            "organization": self._mark_confidence_level("organization", confidence_levels),
            "field": self._mark_confidence_level("field", confidence_levels),
            "talent_rank": self.talent_rank
        }

    def _mark_confidence_level(self, field, confidence_levels):
        """
        根据置信度等级系统，将字段的置信度映射为相应描述。
        """
        confidence = self.confidence_scores.get(field, 0)
        for level_name, (min_conf, max_conf) in confidence_levels.items():
            if min_conf <= confidence <= max_conf:
                return self.__dict__[field] if confidence >= 0.3 else "N/A"  # 0.3为显示信息的最低置信度
        return "N/A"

# 置信度等级映射词典
CONFIDENCE_LEVELS = {
    "Very High": (0.9, 1.0),
    "High": (0.7, 0.9),
    "Moderate": (0.5, 0.7),
    "Low": (0.3, 0.5),
    "Very Low": (0.0, 0.3)
}

# API 路由：获取开发者数据，应用置信度等级
@app.route('/get_developers', methods=['GET'])
def get_developers():
    try:
        with neo4j_driver.session() as session:
            # 从数据库获取开发者信息以及置信度数据
            results = session.run("""
                MATCH (d:Developer)
                RETURN d.id AS id, d.name AS name, d.followers AS followers,
                       d.nation AS nation, d.organization AS organization,
                       d.field AS field, d.talent_rank AS talent_rank,
                       d.nation_confidence AS nation_confidence,
                       d.organization_confidence AS organization_confidence,
                       d.field_confidence AS field_confidence
            """)

            developers = []
            for record in results:
                # 提取每个字段的置信度
                confidence_scores = {
                    "nation": record["nation_confidence"],
                    "organization": record["organization_confidence"],
                    "field": record["field_confidence"]
                }

                developer = Developer(
                    id=record["id"],
                    name=record["name"],
                    followers=record["followers"],
                    nation=record["nation"],
                    organization=record["organization"],
                    field=record["field"],
                    talent_rank=record["talent_rank"],
                    confidence_scores=confidence_scores
                )

                # 将开发者数据附加到返回列表，包含置信度等级
                developers.append(developer.with_confidence_levels(CONFIDENCE_LEVELS))

        return jsonify(developers)
    except Exception as e:
        logger.error(f"获取开发者数据失败: {e}")
        return jsonify({"error": "An error occurred while retrieving developer data."}), 500

if __name__ == "__main__":
    app.run(port=5000)