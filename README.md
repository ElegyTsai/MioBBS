# 「 Mio　BBS 」开发手册
本论坛主要参考米游社，作为后端开发学习的练手项目。
## 所用工具
- 语言： Go
- web框架：GIN
- Database: MySQL
- ORM： GORM
- 工具：Swagger，Zap，Viper, Casbin, JWT
- 可能选用的工具: kafka，Redis，ORM，

## 项目目标
### 用户管理模块
- 多设备同时登陆
- 自动续期（token？
- 注册
- 用户管理（可选
### 内容管理模块
- 发布帖子
- 帖子显示
- 帖子搜索（可选

## 设计文档

- [学习笔记记录](note/note.md)

```

    ├─src         （源码）
      ├─router         （路由）
      ├─config         （配置包）
      ├─model          （结构体层）
      ├─service         (服务)
      ├─source         (初始化需要的数据)
      ├─initialize    （初始化）
      ├─log            （一些网站记录文件）
      └─utils          （公共功能）
    ├─note         （学习笔记文件夹）
    
```






## 接口文档



## 记录文档