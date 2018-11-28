package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Bookmark struct {
	c *Config
}

func (b *Bookmark) Add(title string, path string) error {
	bookmarks, _ := b.Get()
	bookmarks[title] = path
	f, err := os.OpenFile(b.c.GetBookmarkPath(), os.O_CREATE|os.O_WRONLY, 0775)
	defer f.Close()
	if err != nil {
		return err
	}
	bk, _ := json.Marshal(bookmarks)
	if _, err = f.Write(bk); err != nil {
		panic(err)
	}
	return nil
}

func (b *Bookmark) Get() (map[string]string, error) {
	f, err := ioutil.ReadFile(b.c.GetBookmarkPath())
	if err != nil {
		return nil, err
	}

	bookmarks := make(map[string]string)
	err = json.Unmarshal(f, &bookmarks)
	if err != nil {
		return make(map[string]string), fmt.Errorf("Bookmark is empty")
	}
	return bookmarks, nil
}
