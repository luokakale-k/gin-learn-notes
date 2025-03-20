# 🚀 Gin Learn Notes

这是我学习 Go Web 框架 [Gin](https://github.com/gin-gonic/gin) 的配套代码仓库，记录从零上手 Gin，到逐步掌握各个功能模块，再到整合开发后端接口项目的全过程。

📚 搭配系列文章发布于掘金 👉 [Gin 实战学习系列](https://juejin.cn/user/888839672963544)

---

## 🔧 功能模块目录

| 章节编号 | 模块内容                     | 状态 |
|------|--------------------------|------|
| 01   | 项目初始化 + 路由拆分             | ✅ 已完成 |
| 02   | POST 接口 + JSON 参数绑定      | 🔜 开发中 |
| 03   | 参数校验（Validator）          | 🔜 开发中 |
| 04   | 集成 GORM + MySQL          | 🔜 开发中 |
| 05   | JWT 登录鉴权                 | 🔜 开发中 |
| 06   | 控制器拆分（request / service） | 🔜 开发中 |
| 07   | 配置文件加载 + 日志中间件           | 🔜 开发中 |
| 08   | 通用分页封装 + 列表接口            | 🔜 开发中 |
| 09   | 统一响应结构 + 错误码处理           | 🔜 开发中 |
| 10   | 集成 Redis 缓存              | 🔜 开发中 |
| 11   | ...............          | 🔜 开发中 |

> 每个子模块都包含对应的示例代码和说明文档，持续更新中…

---

## 🏗 项目结构示意（按功能模块组织）

```
gin-learn-notes/ 
├── config/ # 配置文件读取逻辑 
├── controller/ # 控制器，处理请求入口 
├── core/ # 核心初始化：Gin 引擎、配置加载等
├── logger/ # 日志组件封装 
├── logs/ # 日志输出目录 
├── middleware/ # 中间件（如 JWT、日志）
├── model/ # 数据模型（对应数据库结构）
├── request/ # 请求参数结构体 
├── response/ # 统一响应结构体（建议加） 
├── router/ # 路由注册逻辑 
├── service/ # 服务层，处理业务逻辑 
├── utils/ # 工具函数 
├── vo/ # View Object，展示层数据封装
├── config.yaml # 项目配置文件
├── .gitignore
├── go.mod
└── main.go # 入口文件
```


## 🚀 快速启动

```bash
git clone https://github.com/luokakale-k/gin-learn-notes.git
cd gin-learn-notes

go mod tidy
go run main.go
```

## 📢 项目说明
所有功能按模块维护于主分支中，文章与代码一一对应

配置文件位于 config/config.yaml，请根据实际修改数据库 / Redis 配置

日志默认输出至 logger/logs/ 目录

项目代码简洁清晰，适合学习与二次开发