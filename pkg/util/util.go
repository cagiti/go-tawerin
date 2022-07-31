package util

import (
	"os"
	"path/filepath"
)

const (
	xdgDataFallback = ".local/share"
	tawerinDir      = "tawerin"
	staticDir       = "static"
	templatesDir    = "templates"
)

func HomeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return ""
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func StaticDir() string {
	xdgDataHome := getEnv("XDG_DATA_HOME", filepath.Join(HomeDir(), xdgDataFallback))
	d := filepath.Join(xdgDataHome, tawerinDir, staticDir)
	return d
}

func TemplatesDir() string {
	xdgDataHome := getEnv("XDG_DATA_HOME", filepath.Join(HomeDir(), xdgDataFallback))
	d := filepath.Join(xdgDataHome, tawerinDir, templatesDir)
	return d
}
