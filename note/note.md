[toc]



# 学习笔记记录

## 文档记录相关

- MarkDown用于排版笔记

- 所使用的编辑器是Typora     

  ### [Typora 教程](https://zhuanlan.zhihu.com/p/293557841)

  常用到的快捷键

  - **cmd+1234**      各级标题
  ### MarkDown
  - 表格用| |--来记录

## 项目相关

### 项目文件组织

参考下面的文件组织和[这个笔记](https://www.jianshu.com/p/92919004293d)

    ├─server         （后端文件夹）
    │  ├─api            （API）
    │  ├─config         （配置包）
    │  ├─core           （核心文件）
    │  ├─docs           （swagger文档目录）
    │  ├─global         （全局对象）
    │  ├─initialiaze    （初始化）
    │  ├─middleware     （中间件）
    │  ├─model          （结构体层）
    │  ├─resource       （资源）
    │  ├─router         （路由）
    │  ├─service         (服务)
    │  ├─source         (初始化需要的数据)
    │  ├─plugin         (插件)
    │  └─utils          （公共功能）
    └─web            （前端文件）
        ├─public        （发布模板）
        └─src           （源码包）
            ├─api       （向后台发送ajax的封装层）
            ├─core       （用来修改系统基础可运行配置）
            ├─assets    （静态文件）
            ├─components（组件）
            ├─router    （前端路由）
            ├─store     （vuex 状态管理仓）
            ├─style     （通用样式文件）
            ├─utils     （前端工具库）
            └─view      （前端页面）

