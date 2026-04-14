# DoDoList

DoDoList 是一个用于记录待办事项的 Go 命令行工具，使用 Cobra 搭建喵~

## 可用命令

```bash
dodolist version
dodolist --help
dodolist todo 吃了个汉堡 --pro=2
dodolist list
dodolist list --pro=2
dodolist list --sort
dodolist ok 1
dodolist edit 1 吃了两个汉堡 --pro=3
```

## 目录结构

- `cmd` 命令层，负责各个子命令喵~
- `i18n` 国际化资源加载喵~
- `i18n/locales` 内置语言文件喵~
- `storage` 持久化与日期解析喵~
- `test` 测试目录喵~
- `utils` 命令参数与日期工具喵~

## 说明

所有待办都会保存在程序目录下的 `data/todos.json` 单文件里喵~
`list` 默认按创建时间排序，也支持 `--sort` 按优先级从高到低排序喵~
