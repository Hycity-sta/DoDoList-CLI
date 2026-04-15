# 功能测试报告（英文）

## 测试环境
- 平台：Windows 11 PowerShell
- 项目目录：`c:\Users\Hycity\Documents\cli\DoDoList`
- 构建目录：`test/env`
- 语言环境：`en`
- 报告时间：`2026-04-15`

## 构建结果
- 命令：`go build -a -o .\\test\\env\\dodolist.exe .`
- 结果：成功

## 1. 查看初始空列表
- 命令：`$ .\\dodolist.exe`
- 退出码：`0`
- 结果：成功
```text
INDEX  CREATED AT  STATUS  TODO
-----  ----------  ------  ----
```

## 2. 创建第一个待办
- 命令：`$ .\\dodolist.exe Buy milk`
- 退出码：`0`
- 结果：成功
```text
created todo: Buy milk
```

## 3. 创建第二个待办
- 命令：`$ .\\dodolist.exe Write report`
- 退出码：`0`
- 结果：成功
```text
created todo: Write report
```

## 4. 查看创建后的列表
- 命令：`$ .\\dodolist.exe`
- 退出码：`0`
- 结果：成功
```text
INDEX  CREATED AT           STATUS  TODO        
-----  -------------------  ------  ------------
1      2026-04-15 10:11:52          Buy milk    
2      2026-04-15 10:11:53          Write report
```

## 5. 完成第 1 项待办
- 命令：`$ .\\dodolist.exe ok 1`
- 退出码：`0`
- 结果：成功
```text
completed todo 1
```

## 6. 查看完成后的列表
- 命令：`$ .\\dodolist.exe`
- 退出码：`0`
- 结果：成功
```text
INDEX  CREATED AT           STATUS  TODO        
-----  -------------------  ------  ------------
1      2026-04-15 10:11:52  done    Buy milk    
2      2026-04-15 10:11:53          Write report
```

## 7. 删除第 2 项待办
- 命令：`$ .\\dodolist.exe delete 2`
- 退出码：`0`
- 结果：成功
```text
deleted todo 2
```

## 8. 查看删除后的列表
- 命令：`$ .\\dodolist.exe`
- 退出码：`0`
- 结果：成功
```text
INDEX  CREATED AT           STATUS  TODO    
-----  -------------------  ------  --------
1      2026-04-15 10:11:52  done    Buy milk
```

## 9. 创建第三个待办
- 命令：`$ .\\dodolist.exe Read book`
- 退出码：`0`
- 结果：成功
```text
created todo: Read book
```

## 10. 清除所有已完成待办
- 命令：`$ .\\dodolist.exe clear`
- 退出码：`0`
- 结果：成功
```text
cleared 1 completed todo items
```

## 11. 查看清理后的列表
- 命令：`$ .\\dodolist.exe`
- 退出码：`0`
- 结果：成功
```text
INDEX  CREATED AT           STATUS  TODO     
-----  -------------------  ------  ---------
1      2026-04-15 10:11:54          Read book
```

## 12. 查看当前语言
- 命令：`$ .\\dodolist.exe lang`
- 退出码：`0`
- 结果：成功
```text
current language: en
```

## 13. 查看帮助信息
- 命令：`$ .\\dodolist.exe help`
- 退出码：`0`
- 结果：成功
```text
DoDoList stores todo items in a single local JSON file.

Usage:
  dodolist
  dodolist [command]

Examples:
  dodolist
  dodolist Buy milk
  dodolist ok 1
  dodolist clear
  dodolist delete 1

Available Commands:
  clear       Clear all completed todo items.
  delete      Delete a todo item.
  help        Show help information.
  lang        Set or show language.
  ok          Mark a todo item as completed.
  version     Print version information.

Use "dodolist [command] --help" for more information about a command.
```

## 14. 查看版本信息
- 命令：`$ .\\dodolist.exe version`
- 退出码：`0`
- 结果：成功
```text
dodolist 1.0
```

## 15. 使用不存在的索引完成待办
- 命令：`$ .\\dodolist.exe ok 99`
- 退出码：`1`
- 结果：失败
```text
Error: todo 99 does not exist
Usage:
  dodolist ok [index] [flags]

Flags:
  -h, --help   help for ok
```

## 16. 使用非法索引删除待办
- 命令：`$ .\\dodolist.exe delete 0`
- 退出码：`1`
- 结果：失败
```text
Error: index must be greater than 0
Usage:
  dodolist delete [index] [flags]

Flags:
  -h, --help   help for delete
```

## 17. 切换到不支持的语言
- 命令：`$ .\\dodolist.exe lang jp`
- 退出码：`1`
- 结果：失败
```text
Error: unknown language: jp
Usage:
  dodolist lang [en|zh] [flags]

Flags:
  -h, --help   help for lang
```

## 结论
- 英文环境下的核心功能与错误场景均已实际执行
- 当前测试在 `test/env` 隔离环境中完成
