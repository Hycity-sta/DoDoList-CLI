# DoDoList

DoDoList 是一个用于记录和管理待办事项的命令行工具，基于Cobra进行构建。

## 构建运行

```bash
go mod tidy
go build -o ./bin/dodolist.exe .
cd bin
dodolist version
```

## 可用命令示例

`dodolist`
> 显示所有的待办事项

`dodolist [content]`
> 添加一个待办事项

`dodolist help`
> 显示帮助文件

`dodolist version`
> 显示版本信息

`dodolist ok [content-index]`
> 完成一个待办事项

`dodolist clear`
> 清除所有已完成的待办事项

`dodolist delete [content-index]`
> 删除一个待办事项

`dodolist lang [language]`
> 切换语言

## 平台支持
- windows 11 powershell
