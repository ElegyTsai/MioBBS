# 用户管理设计

基本架构:使用RBAC框架 [可见](https://zhuanlan.zhihu.com/p/63769951)

## 基本模型

User ----->. Role ---> Permission

- User相对于Role 是多对多关系， 即每个User可以对应多个Role

- Role 和Permission也是多对多的关系，每个角色拥有多种权限

## 具体细节

### 基本体系设计相关

#### 角色分类


| 角色名 | 含义 | 详解 |
| ------ | ------ | ------ |
| guest | 未登录用户 | 理论上来说只有临时用户会被划分于此 |
| lv1-6 | 不同论坛等级 | 不同用户可能拥有不同的权限 |
| banned | 被封禁 | 做了什么坏事被关小黑屋 |
| admin | 论坛管理员 | 对论坛内容可以进行管理 |
| system | 系统 | 最高权限，可以管理用户和管理内容 |

但是因为随着论坛功能的不断推新，需要对角色也进行管理，需要对角色有增加删除的功能

#### 权限分类

对帖子的，查看，增加，修改，删除

对用户的，CRUD

对角色的，CRUD



### 数据库和字段设计

#### 用户

- 用户名，可修改

- ***UID***，自动生成，10位，用数据库的自增主键即可

- 邮箱（注册凭证

- 手机号码（我觉得这个有点麻烦，可能只是记录不用手机号码注册

- 密码，需要单向加密？

- 注册时间

#### ~~角色~~

- ~~角色名称~~

- ~~***角色ID*** 主键~~

- ~~角色说明~~

#### ~~权限~~

- ~~**权限ID** 主键~~
- ~~URL地址~~
- ~~访问方式~~

~~**三张实体表格，用于存储固定信息**~~

~~**还有两张关联表格，用来保存彼此之间的映射关系**~~

#### ~~用户-角色~~

~~用户ID，角色ID， 两者一起成为主键，联合索引~~

#### ~~角色-权限~~

~~角色ID，权限ID，同上~~

## Casbin

- P, Role,URL,Method
- G,UID,Role
- 两种策略分别对应 权限和角色，我觉得用角色去分即可。用户只拥有角色，不直接拥有权限，权限由角色去拥有
- 由GORM去修改策略，
