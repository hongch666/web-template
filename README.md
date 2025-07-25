# Demo

本项目是一个基于 Go 语言的 Gin Web 框架和 GORM ORM 的分层示例项目。

## 目录结构

- controller/ 控制器层
- service/ 业务逻辑层
- mapper/ 数据访问层
- model/ 数据模型
- config/ 配置管理
- router/ 路由注册
- main.go 启动入口

## 依赖安装

```sh
go mod tidy
```

## 数据库配置

请在 `config/config.go` 中修改 MySQL 连接参数。

## 启动项目

```sh
go run main.go
```

## 说明

- 默认使用 MySQL 作为数据库。
- 代码结构清晰，便于扩展和维护。
