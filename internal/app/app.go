package app

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/options"
)

// App struct
type App struct {
	Ctx         context.Context
	Settings    *Settings
	WailsConfig *options.App
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Shutdown(ctx context.Context) {
}

type Settings struct {
	Initialized      bool   `json:"Initialized"`
	DefaultAssetsDir string `json:"DefaultAssetsDir"`
	ConfigPath       string `json:"ConfigPath"`
	LoggingFile      string `json:"LoggingFile"`
}

func (a *App) InitSettings(appData string) {
	a.Settings = &Settings{}

	a.Settings.DefaultAssetsDir = "assetsData"
	a.Settings.ConfigPath = filepath.Join(appData, "settings.json")
	a.Settings.Initialized = false
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
	if configFile, err := os.ReadFile(configFile); err != nil {
		return err
	} else {
		if err := json.Unmarshal(configFile, a.Settings); err != nil {
			return err
		}
		return nil
	}
}
