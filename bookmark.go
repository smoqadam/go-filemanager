package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Bookmark struct {
	path string
}

func NewBookmark(p string) Bookmark {
	return Bookmark{path: p}
}
func (b *Bookmark) Add(title string, path string) error {
	bookmarks, _ := b.Get()
	bookmarks[title] = path
	f, err := os.OpenFile(b.path, os.O_WRONLY, 0775)
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
	f, err := ioutil.ReadFile(b.path)
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
