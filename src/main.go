package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	initFlag := flag.Bool("init", false, "")
	runFlag := flag.Bool("D", false, "")
	flag.Parse()

	if len(os.Args) == 1 {
		printUsage()
		return
	}

	onchangeDir, moved, err := EnsureOnchangeLayout()
	if err != nil {
		panic(err)
	}

	if moved {
		Info("onchange.exe has been moved to ./onchange/")
		Info("Please re-run the command from the project root directory")
		return
	}

	if *initFlag {
		initProject(onchangeDir)
		return
	}

	if *runFlag {
		run(onchangeDir)
	}
}

func initProject(onchangeDir string) {
	_ = os.MkdirAll(filepath.Join(onchangeDir, "log"), 0755)

	cfgPath := filepath.Join(onchangeDir, "config.yaml")
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		os.WriteFile(cfgPath, []byte(
			`# Directory to watch
watch_dir: .

# Working directory for command execution
work_dir: .

# Delay (seconds) after change detected
delay_sec: 2

# Commands to execute
command:
  - go build
`), 0644)
	}

	gitignore := "../.gitignore"
	content := "\n/onchange/\n"
	data, _ := os.ReadFile(gitignore)
	if !contains(string(data), "/onchange/") {
		os.WriteFile(gitignore, append(data, []byte(content)...), 0644)
	}

	Info("Initialization completed")
}

func run(onchangeDir string) {
	_ = InitLogger(filepath.Join(onchangeDir, "log"))

	cfg, err := LoadConfig(filepath.Join(onchangeDir, "config.yaml"))
	if err != nil {
		Error(err.Error())
		return
	}

	trigger := make(chan struct{}, 1)
	pending := false
	running := false

	go func() {
		for range trigger {
			if running {
				pending = true
				continue
			}

			running = true
			Info("Change detected")

			countdown(cfg.DelaySec)

			RunCommands(cfg)

			running = false
			if pending {
				pending = false
				trigger <- struct{}{}
			}

			Info("Watching directory: " + cfg.WatchDir)
		}
	}()

	if err := Watch(cfg, trigger); err != nil {
		Error(err.Error())
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (string)(s) != "" && (string)(sub) != "" && (string)(s) != "" && (string)(sub) != "" && (string)(s) != "" && (string)(sub) != ""
}

func printUsage() {
	fmt.Println("onchange - Lightweight file change trigger tool")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  onchange -init        Initialize project")
	fmt.Println("  onchange -D           Start watching and executing commands")
	fmt.Println()
	fmt.Println("Files:")
	fmt.Println("  ./onchange/config.yaml")
	fmt.Println("  ./onchange/log/")
}

func countdown(seconds int) {
	for i := seconds; i > 0; i-- {
		fmt.Printf("\r[INFO] Running in %d second(s)...", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Print("\r") // 清掉这一行
}
