# XCel

XCel是一个强大的Excel处理框架，类似于MSFConsole，提供了灵活的命令行界面来处理Excel文件。

## 功能特性

- ✅ Excel转CSV转换功能
- ✅ CSV转Excel转换功能
- ✅ 支持多Sheet处理
- ✅ 支持多个CSV文件合并为一个Excel文件
- ✅ 列统计分析功能，以JSON格式输出值域
- ✅ 支持多种统计类型：set（去重统计）和bucket（区间统计）
- ✅ 支持按列名或列号指定要统计的列
- ✅ 自动编码转换（GBK → UTF-8）
- ✅ 命令行自动补全支持
- ✅ 跨平台兼容
- ✅ head命令：查看CSV文件前几行内容
- ✅ tail命令：查看CSV文件后几行内容

## 安装

### 前提条件
- Go 1.20或更高版本

### 从源码构建

# 克隆仓库
git clone <repository-url>
cd XCel

# 构建可执行文件

## Windows
```bash
go build -o xcel.exe cmd/xcel/main.go
```

## Linux/macOS
```bash
go build -o xcel cmd/xcel/main.go
```

## 交叉编译（在Windows上构建其他平台的可执行文件）

### 构建Linux 64位可执行文件
```bash
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o xcel_linux_amd64 cmd/xcel/main.go
```

### 构建macOS 64位可执行文件
```bash
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o xcel_darwin_amd64 cmd/xcel/main.go
```

## 或者直接运行
```bash
go run cmd/xcel/main.go <command>
```

## 使用方法

### 基本命令格式

```bash
xcel [command] [flags]
```

### 转换Excel到CSV

```bash
# 基本转换
xcel convert input.xlsx

# 指定输出格式
xcel convert input.xlsx -f csv
xcel convert input.xlsx -t csv
```

### 将多个CSV转换为Excel

```bash
# 将多个CSV文件转换为一个Excel文件（使用逗号分隔）
xcel convert a.csv,b.csv -f xlsx

# 将多个CSV文件转换为一个Excel文件（使用空格分隔）
xcel convert a.csv b.csv -f xlsx
```

### 查看帮助

```bash
# 查看所有命令
xcel help

# 查看特定命令帮助
xcel convert --help
```

## 命令说明

### convert

在Excel和CSV格式之间转换文件。
**用法：**
```bash
xcel convert [input-files...] [flags]
```

**参数：**
- `input-files`：要转换的文件路径，支持多个文件（例如：file1.csv,file2.csv 或 file1.csv file2.csv）

**选项：**
- `-f, --format string`：指定输出格式（默认：csv）
- `-t, --type string`：指定输出格式（默认：csv）（与-f/--format功能相同）
- `-h, --help`：显示命令帮助

### col_stat

按列分析CSV文件，支持多种统计类型，并以JSON格式输出结果。
**用法：**
```bash
xcel col_stat [input-file] [flags]
```

**参数：**
- `input-file`：要分析的CSV文件路径

**选项：**
- `-t, --type string`：指定统计类型（默认：set）
  - `set`：去重统计，输出唯一值列表和计数
  - `bucket`：区间统计，将数值分为多个区间并统计每个区间的数量
- `-n, --number int`：指定bucket统计的桶数量（默认：10）
- `-c, --column string`：指定要统计的列，可以是列名（如：'性别'）或列号（从1开始，如：3）
- `-h, --help`：显示命令帮助

**示例：**
```bash
# 使用默认的set统计类型分析所有列
xcel col_stat data/input/csv/Sheet1.csv

# 使用bucket统计类型，将数值分为20个区间
xcel col_stat data/input/csv/Sheet1.csv -t bucket -n 20

# 只统计第3列
xcel col_stat data/input/csv/Sheet1.csv -c 3

# 只统计名为"性别"的列
xcel col_stat data/input/csv/Sheet1.csv -c '性别'
```

### head

查看CSV文件的前几行内容，包括标题行。
**用法：**
```bash
xcel head [file] [flags]
```

**参数：**
- `file`：要查看的CSV文件路径

**选项：**
- `-n, --lines int`：指定要显示的行数（默认：5）
- `-h, --help`：显示命令帮助

**示例：**
```bash
# 查看文件前5行（默认）
xcel head test.csv

# 查看文件前10行
xcel head test.csv -n 10
```

### tail

查看CSV文件的最后几行内容，包括标题行。
**用法：**
```bash
xcel tail [file] [flags]
```

**参数：**
- `file`：要查看的CSV文件路径

**选项：**
- `-n, --lines int`：指定要显示的行数（默认：10）
- `-h, --help`：显示命令帮助

**示例：**
```bash
# 查看文件最后10行（默认）
xcel tail test.csv

# 查看文件最后15行
xcel tail test.csv -n 15
```

## 版本历史

### v0.0.5 (2025-12-19)

**主要变化：** 新增文件查看命令

- 新增head命令，用于查看CSV文件前几行内容
- 新增tail命令，用于查看CSV文件最后几行内容
- 两个命令都支持-n/--lines选项自定义显示行数
- 确保标题行始终显示
- head命令默认显示前5行
- tail命令默认显示最后10行

### v0.0.4 (2025-12-18)

**主要变化：** 改进列统计分析功能

- 将column_statistic命令重命名为col_stat，更简短易用
- 新增统计类型参数-t/--type，支持set和bucket两种统计类型
- 新增桶数量参数-n/--number，用于bucket统计
- 新增列指定参数-c/--column，支持按列名或列号指定要统计的列
- 优化了统计结果的JSON格式，添加了统计类型标识
- 改进了set统计的唯一值排序
- 实现了bucket统计的区间计算和分布统计

### v0.0.3 (2025-12-18)

**主要变化：** 新增列统计分析功能

- 新增column_statistic命令
- 支持按列分析CSV文件
- 计算每列的值域、唯一值和计数
- 以JSON格式输出统计结果
- 自动保存统计结果到JSON文件

### v0.0.2 (2025-12-18)

**主要变化：** 新增CSV到Excel转换功能

- 支持将多个CSV文件转换为一个Excel文件
- 支持使用逗号或空格分隔多个输入文件
- 新增-f/--format参数用于指定输出格式
- 保持与原有-t/--type参数的兼容性

### v0.0.1 (2025-12-18)

**主要变化：** 搭建项目的基础框架

- 初始化Go模块
- 搭建基于Cobra的命令行框架
- 实现convert命令的基础结构
- 集成excelize库用于Excel处理
- 实现Excel到CSV的转换功能
- 支持多Sheet处理
- 支持编码转换

## 开发计划

### 短期计划（1-2个月）
- [ ] 添加更多输出格式支持（JSON、TXT等）
- [ ] 支持Sheet筛选功能
- [ ] 优化性能和内存占用
- [ ] 添加单元测试
- [ ] 支持数据排序和过滤功能
- [ ] 增强列统计分析功能，支持更复杂的统计指标

### 中期计划（3-6个月）
- [ ] 实现批量操作功能，支持对多个文件执行相同命令
- [ ] 添加命令行交互模式，支持连续操作
- [ ] 实现数据查找和替换功能
- [ ] 支持密码保护的Excel文件处理
- [ ] 添加配置文件支持，允许用户保存常用配置

### 长期计划（6个月以上）
- [ ] 实现图表生成功能
- [ ] 支持Excel公式计算
- [ ] 实现插件系统，允许扩展功能
- [ ] 支持数据导入导出到数据库
- [ ] 添加可视化报告生成功能
- [ ] 完善文档和教程
- [ ] 增加更多统计分析算法
- [ ] 支持更多文件格式（PDF、HTML等）

## 项目结构

```
XCel/
├── cmd/
│   └── xcel/
│       └── main.go          # 主入口文件
├── internal/
│   ├── commands/            # 命令实现
│   │   └── convert.go       # convert命令实现
│   └── utils/               # 工具函数
│       └── excel.go         # Excel处理工具
├── go.mod                   # Go模块文件
├── go.sum                   # 依赖哈希值文件
└── README.md                # 项目说明文档
```

## 许可证

MIT License

## 贡献

欢迎提交Issue和Pull Request！

## 联系方式

如有问题或建议，请通过GitHub Issues反馈。