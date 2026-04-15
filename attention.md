# 开发注意事项

## 目录结构

```text
DoDoList
├─ cmd 存放命令
├─ storage 持久化配置
├─ utils 工具集
├─ config 配置文件
├─ i18n 国际化支持
│  └─ locales
├─ bin 构建存放目录
│  └─ data 持久化数据存放目录
└─ .vscode vscode工作区
```

## 注释规范

- 不使用godoc规范
- 在每个方法前都注释用来干什么的
- 在流程中进行注释
- 使用中文注释

## 代码规范

- 流程与流程中用空行隔开

## 功能测试
- 将程序构建在测试环境test/env下
- 用ai将所有功能都测试一遍，并将测试用例与测试结果输出到test/output.md文件下
- 测试要分语言，不同语言的测试结果存放在output-zh.md类似的语言输出中

## 单元测试
- 用ai生成单元测试, 放置在test/unit下
- 使用go test来执行单元测试
