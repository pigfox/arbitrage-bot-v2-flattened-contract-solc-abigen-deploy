package util

import (
	"log"
	"os"
	"path/filepath"
)

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error:", err)
		return ""
	}
	return dir
}

func EmptyFolder(folderPath string) error {
	d, err := os.Open(folderPath)
	if err != nil {
		return err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		fullPath := filepath.Join(folderPath, name)
		err := os.RemoveAll(fullPath)
		if err != nil {
			return err
		}
	}
	return nil
}
