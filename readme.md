# DoDoList

DoDoList 是一个用于记录待办事项的 Go 命令行工具，使用 Cobra 搭建喵~

## 构建运行

```bash
go build -o ./bin/dodolist.exe .
./bin/dodolist.exe help
```

## 可用命令示例

```bash
dodolist version
dodolist help

dodolist todo 学习 Go Cobra --pro=3
dodolist todo 写项目文档 --pro=1
dodolist todo 修复一个小虫子 --pro=2

dodolist show
dodolist show --pro=2
dodolist show --sort
dodolist show --status-sort

dodolist ok 2
dodolist edit 1 学习 Go Cobra 和文件存储 --pro=4
dodolist delete 3
dodolist show --sort --status-sort
```

## 目录结构

```text
DoDoList
├─ cmd
├─ storage
├─ utils
├─ config
├─ i18n
│  └─ locales
├─ bin
│  └─ data
└─ .vscode
```

- `cmd` 负责命令定义和执行流程喵~
- `storage` 负责待办数据的单文件持久化喵~
- `utils` 放命令共用的小工具函数喵~
- `config` 和 `i18n` 预留给语言配置与国际化能力喵~
- `config.json` 会保存在可执行文件同级目录，用来保存当前界面语言配置喵~
- `bin/data/todos.json` 是运行后生成的待办数据文件喵~
