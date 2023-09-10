package app

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// App struct
type App struct {
	ctx         context.Context
	settings    *Settings
	wailsConfig *options.App
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) shutdown(ctx context.Context) {
}

type Settings struct {
	Initialized      bool   `json:"initialized"`
	DefaultAssetsDir string `json:"defaultAssetsDir"`
	ConfigPath       string `json:"configPath"`
}

func Init(assetDir string) *Settings {
	return &Settings{
		DefaultAssetsDir: assetDir,
		ConfigPath:       "./settings.json",
		Initialized:      false,
	}
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
