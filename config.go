package main

import (
	"os"
)

type Config struct {
	dirPath      string
	filePath     string
	bookmarkPath string
}

func NewConfig() *Config {

	p := UserHomeDir() + "/.gofila"
	c := &Config{}
	c.dirPath = p
	c.createDir()
	c.creteFile("config")
	c.creteFile("bookmarks")
	c.bookmarkPath = p + "/bookmarks"

	return c
}

func (c *Config) GetBookmarkPath() string {
	return c.bookmarkPath
}

func (c *Config) createDir() {
	if _, err := os.Stat(c.dirPath); os.IsNotExist(err) {
		os.Mkdir(c.dirPath, 775)
	}
}

func (c *Config) creteFile(p string) string {
	_, err := os.OpenFile(c.dirPath+"/"+p, os.O_CREATE|os.O_WRONLY, 0775)
	if err != nil {
		return err.Error()
	}
	return c.dirPath + "/" + p
}

func UserHomeDir() string {
	return os.Getenv("HOME")
}

func (c *Config) GetDirPath() string {
	return c.dirPath
}
