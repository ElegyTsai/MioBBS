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
  - ```-[ ]```可以用来表示勾选

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



##  ZAP

zap是一个高效的logger记录，可参考[blog](https://www.liwenzhou.com/posts/Go/zap/)

在相关的项目中，有一种比较好的方式是

重写一个printf函数，把zap包装一下，因为有一些记录可能会记录进mysql，通过包装可以更方便的使用

## VIPER

[Viper GitHub](https://github.com/spf13/viper)

viper 是一个配置文件管理的包，主要逻辑是从yaml等文件中读取配置，方便我们修改。[一个简单的教程](https://www.liwenzhou.com/posts/Go/viper_tutorial/)

1. 在指定文件夹下建立一个yaml文件
2. 用viper 去读入这些配置文件
3. 在代码中使用对应的键值，完成配置

## Gorm

[官网](https://gorm.io/zh_CN/docs/connecting_to_the_database.html)

用Viper去保存位置






## Go语言相关

- 用 . import 可以省略前面的
- if := 中声明的变量会在if语句块结束后消失
- *_test.go 文件用于默认的测试框架
- 注意 go test/run/build 所在的文件都是不一样的 
