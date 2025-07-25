# Demo

本项目基于 [NestJS](https://nestjs.com/) 框架，集成 TypeORM，演示如何查询 MySQL 数据库中的用户（id 和 name 字段）。

## 快速开始

### 1. 安装依赖

```bash
npm install
```

### 2. 配置数据库

在 `src/app.module.ts` 中修改 TypeORM 配置，填写你的 MySQL 连接信息：

```
TypeOrmModule.forRoot({
  type: 'mysql',
  host: 'localhost',
  port: 3306,
  username: 'root', // 修改为你的用户名
  password: 'your_password', // 修改为你的密码
  database: 'your_database', // 修改为你的数据库名
  autoLoadEntities: true,
  synchronize: true, // 生产环境建议关闭
}),
```

### 3. 启动项目

```bash
npm run start:dev
```

### 4. 访问用户接口

- 查询所有用户：
  - GET `http://localhost:8080/users`

## 目录结构

```
src/
  user/
    user.entity.ts      # 用户实体
    user.service.ts     # 用户服务
    user.controller.ts  # 用户控制器
    user.module.ts      # 用户模块
  app.module.ts         # 主模块
```

## 其他

- 推荐使用 `.env` 文件管理数据库等敏感配置。
- 如需扩展用户表字段，请在 `user.entity.ts` 中修改。
