[中文版本](#-功能特性)
---

# onChange

`onChange` is a **directory change monitoring tool** designed to **execute a set of custom commands with a delay** whenever files in a specified directory change. It is suitable for scenarios such as automatic building, deployment, synchronization, development helper scripts, and more.

---

## ✨ Features

* 📂 Monitors file changes in a specified directory (using `fsnotify`)
* ⏱ Supports **delayed execution** (debouncing to prevent frequent triggers)
* 🔁 Automatically merges events during a change cycle — **executes only once per cycle**
* 🧾 Supports configuration file `config.yaml`
* 📜 Automatically generates logs
* 🧠 If changes occur while running, another execution will occur after the current one finishes

---

## 📁 Directory Structure

```text
onChange/
├── onChange.exe        # Main program (Windows)
├── config.yaml         # Configuration file
├── log/                # Log directory (created automatically)
└── README.md
```

---

## ⚙️ Configuration File (config.yaml)

Example:

```yaml
watch_dir: .
delay_sec: 5
commands:
  - echo "File changed"
  - go build
```

Field Description:

| Field          | Description                           |
| -------------- | ------------------------------------- |
| `watch_dir`    | Directory path to monitor             |
| `delay_sec`    | Seconds to delay execution after a trigger (debounce) |
| `commands`     | List of commands to execute after a change is detected |

---

## 🚀 Usage

### 1️⃣ Run Directly

```bash
onChange
```

* Automatically loads `config.yaml` if it exists in the current directory
* Shows **usage instructions** when no arguments are provided

---

### 2️⃣ Initialize (init)

Run inside the `onChange` directory:

```bash
onChange -init
```

Effects:

* Generates a default `config.yaml`
* Creates the log directory
* **If `init` is executed inside the `onChange/` directory, it will automatically delete the `onChange.exe` in the parent directory to avoid duplication**

---

## ⏳ Change Trigger Logic (Important)

* File change → triggers monitoring
* Starts a countdown (`delay_sec` seconds)
* During the countdown:

  * If another change occurs, it will **not** execute again immediately
  * Will **execute once more** after the current execution finishes
* Only **one execution per change cycle**

👉 This design prevents multiple executions triggered by saving files rapidly.

---

## 🧠 Internal Mechanism Overview

* Uses `fsnotify` to monitor the directory
* Uses a buffered `trigger channel` to merge events
* Uses `running / pending` status to avoid duplicate executions
* Main thread keeps the program running via `time.Sleep(365 days)`

---

## 📜 Logs

* Logs are stored by default in:

```text
onChange/log/
```

* Contains:

  * Startup information
  * Change detection events
  * Command execution status
  * Error messages

---

## 📌 Suitable Scenarios

* Automated builds / compilation
* File synchronization triggers
* Development helper scripts
* Simple CI / local automation

---

---

# onChange

`onChange` 是一个**目录变更监听工具**，用于在指定目录发生文件变化时，**延迟执行一组自定义命令**。
适合用于自动构建、自动部署、自动同步、开发辅助脚本等场景。

---

## ✨ 功能特性

* 📂 监听指定目录的文件变化（基于 `fsnotify`）
* ⏱ 支持 **延迟执行**（防抖，避免频繁触发）
* 🔁 变更期间自动合并事件，**只执行一次**
* 🧾 支持配置文件 `config.yaml`
* 📜 自动生成日志
* 🧠 运行中再次变更会在当前执行结束后再执行一次

---

## 📁 目录结构

```text
onChange/
├── onChange.exe        # 主程序（Windows）
├── config.yaml         # 配置文件
├── log/                # 日志目录（自动创建）
└── README.md
```

---

## ⚙️ 配置文件说明（config.yaml）

示例：

```yaml
watch_dir: .
delay_sec: 5
commands:
  - echo "File changed"
  - go build
```

字段说明：

| 字段名         | 说明             |
| ----------- | -------------- |
| `watch_dir` | 要监听的目录路径       |
| `delay_sec` | 触发后延迟执行的秒数（防抖） |
| `commands`  | 发生变化后要执行的命令列表  |

---

## 🚀 使用方法

### 1️⃣ 直接运行

```bash
onChange
```

* 如果当前目录下存在 `config.yaml`，会自动加载
* 未提供参数时会输出**使用说明**

---

### 2️⃣ 初始化（init）

在 `onChange` 目录中运行：

```bash
onChange -init
```

作用：

* 生成默认 `config.yaml`
* 创建日志目录
* **如果在 `onChange/` 目录中执行 `init`，会自动删除上一级的 `onChange.exe`，避免重复存在**

---

## ⏳ 变更触发逻辑说明（重要）

* 文件发生变化 → 触发监听
* 开始倒计时（`delay_sec` 秒）
* 倒计时期间：

  * 如果再次发生变化，不会重复执行
  * 会在当前执行完成后 **再执行一次**
* 同一轮变化 **最多执行一次**

👉 这是为了避免保存文件时触发多次执行的问题。

---

## 🧠 内部机制简述

* 使用 `fsnotify` 监听目录
* 使用带缓冲的 `trigger channel` 合并事件
* 使用 `running / pending` 状态避免重复执行
* 主线程通过 `time.Sleep(365 天)` 保持程序运行

---

## 📜 日志

* 日志默认存放在：

```text
onChange/log/
```

* 包含：

  * 启动信息
  * 变更检测
  * 命令执行状态
  * 错误信息

---


## 📌 适用场景

* 自动构建 / 编译
* 文件同步触发
* 开发过程辅助脚本
* 简单 CI / 本地自动化

---
