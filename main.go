package main

import (
	"fmt"

	"./filemanager"
	"./menu"
	termbox "github.com/nsf/termbox-go"
)

var fm filemanager.FileManager
var bookmark Bookmark
var config Config

const (
	PAGE_FILES     string = "PAGE"
	PAGE_BOOKMARKS string = "BOOKMARK"
	PAGE_CONFIRM   string = "CONFIRM"
)

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

	config := NewConfig()
	bookmark = NewBookmark(config.GetBookmarkPath())

	m.AddEvent(termbox.KeyCtrlB, setBookmark)
	m.AddEvent(termbox.KeyCtrlSlash, showBookmarks)
	m.AddEvent(termbox.KeyEnter, enter)
	m.AddEvent(termbox.KeyArrowRight, enter)
	m.AddEvent(termbox.KeyArrowLeft, back)
	m.AddEvent(termbox.KeyEsc, exit)
	m.AddEvent(termbox.KeyDelete, del)

	m.Render()
}

func refresh(m *menu.Menu, f *filemanager.FileManager, files []filemanager.File, currentPage string) error {
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

	m.SetItems(menuItems, currentPage)
	if currentPage == PAGE_BOOKMARKS {
		m.Info("Bookmarks")
	} else {
		m.Info("Path: " + f.Path().Current())
	}
	return nil
}

func showConfirm(m *menu.Menu, val string) {
	confirmItems := []menu.MenuItem{}
	confirmItems = append(confirmItems, menu.MenuItem{
		Title: "No",
		Value: "no",
	})
	confirmItems = append(confirmItems, menu.MenuItem{
		Title: "Yes",
		Value: val,
	})
	m.Info(fmt.Sprintf("Do you want to delete %s?", val))
	m.SetItems(confirmItems, PAGE_CONFIRM)
}
