package main

import (
	"fmt"
	"os"
)

type Config struct {
	path         string
	bookmarkPath string
}

func NewConfig() *Config {
	p := UserHomeDir() + "/.gofila"
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.Mkdir(p, 0775)
		fmt.Print(p)
	}
	return &Config{
		bookmarkPath: p + "/bookmarks",
	}
}

func (c *Config) GetBookmarkPath() string {
	return c.bookmarkPath
}

func UserHomeDir() string {
	// if runtime.GOOS == "windows" {
	// 	home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
	// 	if home == "" {
	// 		home = os.Getenv("USERPROFILE")
	// 	}
	// 	return home
	// }
	return os.Getenv("HOME")
}
