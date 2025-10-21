package main

import (
	"errors"
	"os"
)

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}

func fileRead(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}

	return string(content)
}
