package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func getThemeDirectory() string {
	theme, ok := os.LookupEnv("THEME_NAME")
	if !ok {
		theme = "default"
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	themeDir := path.Join(cwd, "skins", theme)

	if !fileExists(path.Join(themeDir, "xxx.hbs")) {
		log.Fatalf("theme: %s not found\n", themeDir)
	}

	return themeDir
}

func getThemeViewFile(themeDir string, statusCode int64) string {
	statusCodeStr := fmt.Sprint(statusCode)

	if fileExists(path.Join(themeDir, statusCodeStr+".hbs")) {
		return statusCodeStr
	}

	if fileExists(path.Join(themeDir, statusCodeStr[:1]+"xx.hbs")) {
		return statusCodeStr[:1] + "xx"
	}

	return "xxx"
}
