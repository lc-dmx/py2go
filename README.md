# py2go

目标：寻找一些有意思的 python 包，用 Go 重新实现一下。

## 目录结构及描述

```bash
    ├── packages // 各种 Python 包的实现
    │   └── fake_useragent
    ├── go.mod // 依赖包
    └── README.md
```

每个实现包结构如下：
```
    ├── common // 常量以及配置定义
        ├── const.go
        └── error.go
    ├── compare // 实现效果比较
    ├── core // 核心实现
    ├── main // main 函数 
    ├── test // 测试用例目录 
    └── utils // 工具函数目录
```
