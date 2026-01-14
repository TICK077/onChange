package main

import (
	"os"
	"path/filepath"
)

func EnsureOnchangeLayout() (string, bool, error) {

	cwd, _ := os.Getwd()
	onchangeDir := filepath.Join(cwd, "onchange")

	exe, _ := os.Executable()
	exeBase := filepath.Base(exe)
	exe, _ = filepath.Abs(exe)
	exeName := filepath.Base(exe)

	// already in ./onchange
	if filepath.Base(cwd) == "onchange" {
		parentExe := filepath.Join(filepath.Dir(cwd), exeName)
		_ = os.Remove(parentExe) // ignore error
		return cwd, false, nil
	}

	if _, err := os.Stat(onchangeDir); os.IsNotExist(err) {
		if err := os.Mkdir(onchangeDir, 0755); err != nil {
			return "", false, err
		}
	}

	targetExe := filepath.Join(onchangeDir, exeBase)

	if _, err := os.Stat(targetExe); os.IsNotExist(err) {
		if err := copyFile(exe, targetExe); err != nil {
			return "", false, err
		}
		return onchangeDir, true, nil
	}

	return onchangeDir, false, nil
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0755)
}
