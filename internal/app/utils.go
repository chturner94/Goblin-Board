package app

import (
	"os"
)

type Directory struct {
	absPath string
	subDirs []Directory
}

func (d *Directory) List() []string {
	path := d.absPath
}
