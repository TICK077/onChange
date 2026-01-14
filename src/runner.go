package main

import (
	"os/exec"
)

func RunCommands(cfg *Config) bool {
	for _, cmdStr := range cfg.Command {
		Info("Running: " + cmdStr)

		cmd := exec.Command("cmd", "/C", cmdStr)
		cmd.Dir = cfg.WorkDir

		out, err := cmd.CombinedOutput()
		if len(out) > 0 {
			Logger.Println(string(out))
		}

		if err != nil {
			Error("Command failed: " + cmdStr)
			return false
		}
	}
	return true
}
