# 功能测试报告（中文）

## 测试环境
- 平台：Windows 11 PowerShell
- 项目目录：`c:\Users\Hycity\Documents\cli\DoDoList`
- 构建目录：`test/env`
- 语言环境：`zh`
- 报告时间：`2026-04-15`

## 构建结果
- 命令：`go build -a -o .\\test\\env\\dodolist.exe .`
- 结果：成功

## 1. 查看初始空列表
- 命令：`$ .\\dodolist.exe`
- 退出码：`0`
- 结果：成功
```text
序号  创建时间  状态  待办
----  --------  ----  ----
```

## 2. 创建第一个待办
- 命令：`$ .\\dodolist.exe 买牛奶`
- 退出码：`0`
- 结果：成功
```text
已创建待办：买牛奶
```

## 3. 创建第二个待办
- 命令：`$ .\\dodolist.exe 写周报`
- 退出码：`0`
- 结果：成功
```text
已创建待办：写周报
```

## 4. 查看创建后的列表
- 命令：`$ .\\dodolist.exe`
- 退出码：`0`
- 结果：成功
```text
序号  创建时间             状态  待办  
----  -------------------  ----  ------
1     2026-04-15 10:12:04        买牛奶
2     2026-04-15 10:12:05        写周报
```

## 5. 完成第 1 项待办
- 命令：`$ .\\dodolist.exe ok 1`
- 退出码：`0`
- 结果：成功
```text
已完成待办 1
```

## 6. 查看完成后的列表
- 命令：`$ .\\dodolist.exe`
- 退出码：`0`
- 结果：成功
```text
序号  创建时间             状态    待办  
----  -------------------  ------  ------
1     2026-04-15 10:12:04  已完成  买牛奶
2     2026-04-15 10:12:05          写周报
```

## 7. 删除第 2 项待办
- 命令：`$ .\\dodolist.exe delete 2`
- 退出码：`0`
- 结果：成功
```text
已删除待办 2
```

## 8. 查看删除后的列表
- 命令：`$ .\\dodolist.exe`
- 退出码：`0`
- 结果：成功
```text
序号  创建时间             状态    待办  
----  -------------------  ------  ------
1     2026-04-15 10:12:04  已完成  买牛奶
```

## 9. 创建第三个待办
- 命令：`$ .\\dodolist.exe 读书`
- 退出码：`0`
- 结果：成功
```text
已创建待办：读书
```

## 10. 清除所有已完成待办
- 命令：`$ .\\dodolist.exe clear`
- 退出码：`0`
- 结果：成功
```text
已清除 1 个已完成待办
```

## 11. 查看清理后的列表
- 命令：`$ .\\dodolist.exe`
- 退出码：`0`
- 结果：成功
```text
序号  创建时间             状态  待办
----  -------------------  ----  ----
1     2026-04-15 10:12:06        读书
```

## 12. 查看当前语言
- 命令：`$ .\\dodolist.exe lang`
- 退出码：`0`
- 结果：成功
```text
当前语言：zh
```

## 13. 查看帮助信息
- 命令：`$ .\\dodolist.exe help`
- 退出码：`0`
- 结果：成功
```text
DoDoList 将待办事项存放在一个本地 JSON 文件中。

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
  clear       清除所有已完成的待办事项。
  delete      删除一个待办事项。
  help        显示帮助信息。
  lang        设置或查看语言。
  ok          将待办事项标记为已完成。
  version     输出版本信息。

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
Error: 待办 99 不存在
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
Error: 索引必须大于 0
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
Error: 未知语言：jp
Usage:
  dodolist lang [en|zh] [flags]

Flags:
  -h, --help   help for lang
```

## 结论
- 中文环境下的核心功能与错误场景均已实际执行
- 当前测试在 `test/env` 隔离环境中完成
