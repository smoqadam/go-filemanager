package main

import (
	"fmt"

	"./filemanager"
	"./menu"
	termbox "github.com/nsf/termbox-go"
)

var fm filemanager.FileManager
var b Bookmark

func main() {
	fm = filemanager.New("/")

	menuItems := []menu.MenuItem{}
	files, err := fm.Ls()
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		menuItems = append(menuItems, menu.MenuItem{
			Title: f.Name,
			Value: f.Path,
		})
	}
	m := menu.New(menuItems)
	m.Info("Path: " + fm.Path().Current())

	b.c = NewConfig()
	m.AddEvent(termbox.KeyCtrlB, setBookmark)
	m.AddEvent(termbox.KeyCtrlSlash, showBookmarks)
	m.AddEvent(termbox.KeyEnter, enter)
	m.AddEvent(termbox.KeyArrowRight, enter)
	m.AddEvent(termbox.KeyArrowLeft, back)
	m.AddEvent(termbox.KeyEsc, exit)

	m.Render()
}

func refresh(m *menu.Menu, f *filemanager.FileManager, files []filemanager.File) error {
	menuItems := []menu.MenuItem{}
	if len(files) == 0 {
		return fmt.Errorf("Directory " + f.Path().Current() + " is empty")
	}
	for _, f := range files {
		menuItems = append(menuItems, menu.MenuItem{
			Title: f.Name,
			Value: f.Path,
		})
	}

	m.SetItems(menuItems)
	m.Info("Path: " + f.Path().Current())
	return nil
}
