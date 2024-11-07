# DevInsight

<div align=center>

[![Static Badge](https://img.shields.io/badge/python-3.11.4-blue?style=flat)](https://www.python.org)
[![Static Badge](https://img.shields.io/badge/go--github-v66.0.0-blue)](https://github.com/google/go-github)
[![Static Badge](https://img.shields.io/badge/go--gorm-v1.26.0-red)](https://gorm.io)
[![Static Badge](https://img.shields.io/badge/gin-v1.10.0-8A2BEA?style=flat)](https://pkg.go.dev/github.com/gin-gonic/gin)
[![Static Badge](https://img.shields.io/badge/go-1.22.5-blue?style=flat)](https://pkg.go.dev/golang.org/dl/go1.22.5)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Static Badge](https://img.shields.io/badge/mysql-latest-green?style=flat)](https://www.mysql.com)
[![Static Badge](https://img.shields.io/badge/neo4j-latest-green?style=flat)](https://neo4j.com/)


</div>


### 题目要求

#### 七牛云赛题: github数据应用

> 根据 GitHub 的开源项目数据，开发一款开发者评估应用。

#### 基础功能
- 开发者在技术能力方面 TalentRank（类似 Google 搜索的 PageRank），对开发者的技术能力进行评价/评级。评价/评级依据至少包含：项目的重要程度、该开发者在该项目中的贡献度。
- 开发者的 Nation。有些开发者的 Profile 里面没有写明自己的所属国家/地区。在没有该信息时，可以通过其关系网络猜测其 Nation。
- 开发者的领域。可根据领域搜索匹配，并按 TalentRank 排序。Nation 作为可选的筛选项，比如只需要显示所有位于中国的开发者。
#### 高级功能
- 所有猜测的数据，应该有置信度。置信度低的数据在界面展示为 N/A 值。
- 开发者技术能力评估信息自动整理。有的开发者在 GitHub 上有博客链接，甚至有一些用 GitHub 搭建的网站，也有一些在 GitHub 本身有账号相关介绍。可基于类 ChatGPT 的应用整理出开发者评估信息。

### 代码架构

- [后端代码](/back-end)
  - [go业务部分](/back-end/primary_server)
  - [go数据处理模块](/back-end/data_process)
  - [python算法数据处理模块](/back-end/data_process_algo)
- [前端代码](/front-end)


### 演示视频

[vedio](/assert/vedio.mp4)