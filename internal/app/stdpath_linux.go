package app

import (
	"os"
	"path/filepath"
	"runtime"
)

func appDataLocation(name string) (string, error) {
	userOs := runtime.GOOS
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if userOs == "linux" {
		p := filepath.Join(homeDir, ".local", "share", name)
		if _, err := os.Stat(p); os.IsNotExist(err) {
			if err := os.MkdirAll(p, 0700); err != nil {
				return "", err
			}
		}
		return p, nil
	} else if userOs == "windows" {
		p := filepath.Join(homeDir, "AppData", "Roaming", name)
		if _, err := os.Stat(p); os.IsNotExist(err) {
			if err := os.MkdirAll(p, 0700); err != nil {
				return "", err
			}
		}
		return p, nil
	} else if userOs == "darwin" {

		p := filepath.Join(homeDir, "Library", "Application Support", name)
		if _, err := os.Stat(p); os.IsNotExist(err) {
			if err := os.MkdirAll(p, 0700); err != nil {
				return "", err
			}
		}
		return p, nil
	}
	return "Your OS is not supported.", nil
}
