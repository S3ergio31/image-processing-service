package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func SaveOnDisk(path string, content []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil {
		LogError(err)
		return err
	}

	if err := os.WriteFile(path, content, 0644); err != nil {
		LogError(err)
		return err
	}

	return nil
}

func LogError(err error) {
	if err == nil {
		return
	}

	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		log.Printf("%s error: %s\n", details.Name(), err.Error())
	}
}
