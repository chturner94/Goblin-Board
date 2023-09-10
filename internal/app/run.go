package app

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/wailsapp/wails/v2"
)

var (
	appname = "GomblinBoard"
	assets  embed.FS
)

func Run(app *App) error {
	appData, err := appDataLocation(appname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open add data directory: %v\n", err)
		return err
	}
	defer crashlog(appData)
	wails.Run(app.wailsConfig)
	if err != nil {
		println("Error:", err.Error)
	}
	return nil
}

func crashlog(appData string) {
	if r := recover(); r != nil {
		if _, err := os.Stat(appData); os.IsNotExist(err) {
			os.MkdirAll(appData, 0700)
		}
		var b bytes.Buffer
		b.WriteString(fmt.Sprintf("%+v\n\n", r))
		buf := make([]byte, 1<<20)
		s := runtime.Stack(buf, true)
		b.Write(buf[0:s])
		os.WriteFile(filepath.Join(appData, "crash.log"), b.Bytes(), 0644)
	}
}
