package app

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/options"
	wruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	WailsConfig options.App
	Ctx         context.Context
	Settings    Settings
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx
}

func (a *App) DOMReady(ctx context.Context) {
	wruntime.EventsOn(ctx, "openDirectoryDialog", func(optionalData ...interface{}) {
		a.OpenFileDirectory()
	})
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) OpenFileDirectory() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	fileDialog, err := wruntime.OpenDirectoryDialog(a.Ctx, wruntime.OpenDialogOptions{
		DefaultDirectory: homeDir,
		Title:            "File Directory",
	})
	if err != nil {
		return
	}
	if fileDialog == "" {
		return
	}
	wruntime.EventsEmit(a.Ctx, "openDirectoryDialog_result", fileDialog)
}

/*
func (a *App) Shutdown(ctx context.Context) {
}
*/

type Settings struct {
	DefaultAssetsDir string `json:"DefaultAssetsDir"`
	ConfigPath       string `json:"ConfigPath"`
	LoggingFile      string `json:"LoggingFile"`
	Initialized      bool   `json:"Initialized"`
}

func (a *App) InitSettings(appData string) {
	a.Settings = Settings{}

	a.Settings.DefaultAssetsDir = "assetsData"
	a.Settings.ConfigPath = filepath.Join(appData, "settings.json")
	a.Settings.Initialized = false
	a.Settings.LoggingFile = filepath.Join(appData, "goblin-board.txt")
}

func filepathJoin(appData, s string) {
	panic("unimplemented")
}

func (s *Settings) WriteConfig() error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	fullPath := s.ConfigPath

	err = os.WriteFile(fullPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) loadSettings(appData string) error {
	configFile := filepath.Join(appData, "settings.json")
	data, err := os.ReadFile(configFile)
	if err != nil {
		a.Settings = Settings{}
		return err
	}
	var s Settings
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	a.Settings = s
	return nil
}

/*
func (a *App) loadSettings(appData string) error {
	configFile := filepath.Join(appData, "settings.json")
	if configFile, err := os.ReadFile(configFile); err != nil {
		return err
	} else {
		if err := json.Unmarshal(configFile, a.Settings); err != nil {
			return err
		}
		return nil
	}
}
*/
