package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	b.write(bookmarks)
	return nil
}

func (b *Bookmark) Get() (map[string]string, error) {
	f, err := ioutil.ReadFile(b.path)
	bookmarks := make(map[string]string)
	err = json.Unmarshal(f, &bookmarks)
	if err != nil || len(bookmarks) == 0 {
		return bookmarks, fmt.Errorf("Bookmark is empty")
	}
	return bookmarks, nil
}

func (b *Bookmark) Delete(name string) error {
	bookmarks, err := b.Get()
	if err != nil {
		return err
	}
	delete(bookmarks, name)
	b.write(bookmarks)
	return nil
}

func (b *Bookmark) write(bookmarks map[string]string) error {
	bk, _ := json.Marshal(bookmarks)
	if err := ioutil.WriteFile(b.path, bk, 0775); err != nil {
		return err
	}
	return nil
}
