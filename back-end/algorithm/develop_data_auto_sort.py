import openai
import aiohttp
import asyncio
import hashlib
from bs4 import BeautifulSoup
from flask import Flask, jsonify, request
import logging
import time
from functools import wraps

# OpenAI API 密钥设置
openai.api_key = "密钥"

# 日志记录初始化
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

# Flask 应用初始化
app = Flask(__name__)

# 模拟缓存
CACHE = {}
CACHE_TTL = 3600  # 缓存有效期(秒)

# 限制 API 访问速率的简单实现
LAST_REQUEST_TIME = {}

def rate_limit(max_calls_per_minute=10):
    """为每个开发者 ID 添加简单的限速控制。"""
    def decorator(f):
        @wraps(f)
        def wrapped(*args, **kwargs):
            user_id = request.json.get("dev_profile_url")
            if user_id:
                now = time.time()
                last_call = LAST_REQUEST_TIME.get(user_id, 0)
                if now - last_call < 60 / max_calls_per_minute:
                    return jsonify({"error": "Too many requests. Please try again later."}), 429
                LAST_REQUEST_TIME[user_id] = now
            return f(*args, **kwargs)
        return wrapped
    return decorator

def handle_error(message, status_code=500):
    """统一的错误处理函数"""
    logger.error(message)
    return jsonify({"error": message}), status_code

def generate_cache_key(*args):
    """基于输入参数生成缓存键，用于减少重复 API 调用"""
    return hashlib.md5("".join(args).encode()).hexdigest()

def is_cache_valid(cache_key):
    """检查缓存是否在有效期内。"""
    if cache_key in CACHE and (time.time() - CACHE[cache_key]['timestamp']) < CACHE_TTL:
        return True
    return False

async def fetch_content(url, retry_attempts=3):
    """通用内容抓取方法，包含超时处理、请求头和重试机制"""
    headers = {'User-Agent': 'Mozilla/5.0 (compatible; DeveloperInfoBot/1.0)'}
    async with aiohttp.ClientSession() as session:
        for attempt in range(retry_attempts):
            try:
                async with session.get(url, headers=headers, timeout=5) as response:
                    if response.status == 200:
                        soup = BeautifulSoup(await response.text(), 'html.parser')
                        return " ".join([p.get_text().strip() for p in soup.find_all("p") if p.get_text()])
            except aiohttp.ClientError as e:
                logger.warning(f"Attempt {attempt + 1}: Failed to fetch content from {url} - {e}")
    return ""

class DeveloperInfoExtractor:
    """负责从开发者的 GitHub 简介、博客和 GitHub Pages 中提取信息。"""
    
    def __init__(self, dev_profile_url, blog_url=None, github_pages_url=None):
        self.dev_profile_url = dev_profile_url
        self.blog_url = blog_url
        self.github_pages_url = github_pages_url

    async def collect_all_content(self):
        """异步获取所有内容，并合并为单一字符串。"""
        cache_key = generate_cache_key(self.dev_profile_url, self.blog_url or "", self.github_pages_url or "")
        
        if is_cache_valid(cache_key):
            return CACHE[cache_key]['data']

        # 使用异步抓取内容
        content_tasks = [fetch_content(self.dev_profile_url)]
        if self.blog_url:
            content_tasks.append(fetch_content(self.blog_url))
        if self.github_pages_url:
            content_tasks.append(fetch_content(self.github_pages_url))
        
        content_parts = await asyncio.gather(*content_tasks)
        combined_content = " ".join(content_parts).strip()
        
        # 存储到缓存
        CACHE[cache_key] = {'data': combined_content, 'timestamp': time.time()}
        
        return combined_content

class DeveloperSummaryGenerator:
    """负责使用 GPT 模型生成开发者的技术能力总结。"""
    
    @staticmethod
    def generate_summary(content, max_tokens=300, include_recommendations=True):
        """通过 GPT 模型生成内容总结，包含技术栈、贡献和改进建议。"""
        try:
            prompt = (
                "分析开发者个人资料并创建报告 "
                "报告:\n\n"
                "---核心技术栈:关键语言,框架和技术。\n"
                "---项目贡献:总结显著的贡献和项目类型。\n"
                "---活跃的兴趣领域:特定的感兴趣的领域。\n"
            )
            if include_recommendations:
                prompt += "建议培养的技能:2-3个需要改进的领域。\n"
            prompt += f"\nContent:\n\n{content}\n\n以结构化格式提供摘要."
            
            response = openai.Completion.create(
                model="gpt-3.5-turbo",
                prompt=prompt,
                max_tokens=min(max_tokens, 300),  # 控制 token 数量
                temperature=0.5
            )
            return response.choices[0].text.strip()
        except Exception as e:
            logger.error(f"Failed to generate summary: {e}")
            return "Technical summary could not be generated."

# API 路由：生成开发者技术能力评估
@app.route('/generate_developer_summary', methods=['POST'])
@rate_limit(5)  # 每分钟限制 5 次调用
async def generate_developer_summary():
    data = request.json
    dev_profile_url = data.get("dev_profile_url")
    blog_url = data.get("blog_url")
    github_pages_url = data.get("github_pages_url")
    include_recommendations = data.get("include_recommendations", True)

    if not dev_profile_url:
        return handle_error("Developer profile URL is required.", 400)

    extractor = DeveloperInfoExtractor(dev_profile_url, blog_url, github_pages_url)
    content = await extractor.collect_all_content()

    if not content:
        return handle_error("No content found from provided URLs.", 404)

    summary_generator = DeveloperSummaryGenerator()
    summary = summary_generator.generate_summary(content, include_recommendations=include_recommendations)

    return jsonify({"developer_summary": summary})

if __name__ == "__main__":
    app.run(port=5000)