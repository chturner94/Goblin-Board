package app

import (
	"fmt"
	"os"
	"time"
)

/*
  func Log(message, path string)
  message = string to be logged and should be formated with something like fmt.Sprintf() or equavalent
  path = path to the log file named goblin_log.txt; most often this can be passed from App.Settings.LoggingFile
*/

func Log(message, path string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	log := currentTime + " - " + message + "\n"

	if _, err := f.WriteString(log); err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	if err := f.Close(); err != nil {
		fmt.Println(err)
		return
	}
}
