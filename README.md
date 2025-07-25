# FastAPI 分层 Web 项目

## 项目简介

本项目基于 FastAPI，采用分层架构（controller、service、mapper），通过 SQLAlchemy ORM 操作 MySQL 数据库，使用 uvicorn 启动。

## 目录结构

- main.py
- app/
  - controller/
  - service/
  - mapper/
  - models/
  - config.py

## 启动方式

```bash
python main.py
# 或
uvicorn main:app --reload
```

## 依赖安装

```bash
pip install fastapi uvicorn sqlalchemy pymysql
# 或
pip install -r requirements.txt
```

## 数据库配置

请在 `app/config.py` 中配置 MySQL 连接信息。
