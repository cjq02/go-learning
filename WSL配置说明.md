# Cursor中配置WSL Terminal的说明

## 检测WSL环境

首先检查您的系统中安装了哪些WSL发行版：

```bash
wsl -l -v
```

输出示例：
```
  NAME                   STATE           VERSION
* docker-desktop         Running         2
  Ubuntu-24.04           Running         2
```

## 配置方法

### 方法1：使用默认WSL发行版（推荐）

在 `.vscode/settings.json` 中设置：

```json
{
    "terminal.integrated.shell.windows": "C:\\Windows\\System32\\wsl.exe"
}
```

### 方法2：指定特定的WSL发行版

如果您有多个WSL发行版，可以指定使用特定的一个：

```json
{
    "terminal.integrated.shell.windows": "C:\\Windows\\System32\\wsl.exe",
    "terminal.integrated.shellArgs.windows": ["-d", "Ubuntu-24.04"]
}
```

将 `Ubuntu-24.04` 替换为您实际的发行版名称。

## 测试配置

1. 在Cursor中按 `Ctrl+Shift+P` 打开命令面板
2. 输入 `Terminal: Create New Terminal`
3. 新的terminal应该会自动打开WSL环境

## 故障排除

### 如果terminal没有启动

1. 确保WSL已经正确安装和配置
2. 检查WSL发行版是否正在运行：`wsl -l -v`
3. 如果没有运行，使用：`wsl -d Ubuntu-24.04` 启动

### 如果出现权限错误

确保您有足够的权限运行WSL：

```bash
wsl --shutdown
wsl -d Ubuntu-24.04
```

### 在WSL中安装Go（如果需要）

```bash
# 在WSL中
sudo apt update
sudo apt install golang-go
go version
```

## 高级配置

### 在WSL中运行Go命令

如果您希望Cursor直接在WSL中运行Go命令，可以添加：

```json
{
    "go.alternateTools": {
        "go": "wsl go"
    }
}
```

### 设置默认工作目录

```json
{
    "terminal.integrated.cwd": "${workspaceFolder}"
}
```

这样terminal会自动切换到项目目录。

## 注意事项

- 配置完成后需要重启Cursor才能生效
- 确保您的项目路径在WSL中是可访问的
- 如果遇到路径问题，可以考虑将项目放在WSL文件系统内（`/home/username/projects/`）
