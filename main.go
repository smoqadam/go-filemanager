package main

import (
	"fmt"
	"math"

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

var infoSection menu.Section

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
	m.AddEvent(termbox.KeyArrowDown, GoDown)
	m.AddEvent(termbox.KeyArrowUp, GoUp)
	m.AddEvent(termbox.KeyCtrlI, fileContent)

	infoSection = menu.NewSection(30, 5, m.Window().Width()-45, 0)
	infoSection.Title = " File Info "
	m.AddSection(&infoSection)

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

	// infoSection.SetContent("Information about this file is: ")

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

// stolen from https://github.com/pyk/byten/blob/master/size.go
func index(s int64) float64 {
	x := math.Log(float64(s)) / math.Log(1024)
	return math.Floor(x)
}

func countSize(s int64, i float64) float64 {
	return float64(s) / math.Pow(1024, math.Floor(i))
}

// Size return a formated string from file size
func Size(s int64) string {

	symbols := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	i := index(s)
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}
	size := countSize(s, i)
	format := "%.0f"
	if size < 10 {
		format = "%.1f"
	}

	return fmt.Sprintf(format+"%s", size, symbols[int(i)])
}
