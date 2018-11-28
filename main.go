package main

import (
	"./filemanager"
	"./menu"
	termbox "github.com/nsf/termbox-go"
)
var fm filemanager.FileManager

func main() {
	fm = filemanager.New("/")

	menuItems := []string{}
	files, err := fm.Ls()
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		menuItems = append(menuItems, f.Name)
	}
	m := menu.New(menuItems)
	m.Info("Path: " + fm.Path().Current())

	m.AddEvent(termbox.KeyEnter, enter)
	m.AddEvent(termbox.KeyArrowRight, enter)
	m.AddEvent(termbox.KeyArrowLeft, back)

  m.AddEvent(termbox.KeyEsc, exit)
	m.Render()
}



func refresh(m *menu.Menu, f *filemanager.FileManager, path string) error {
	files, err := f.Ls()
	if err != nil {
		return err
	}
	menuItems := []string{}
	for _, f := range files {
		menuItems = append(menuItems, f.Name)
	}
	e := m.SetStringItems(menuItems)
	if e != nil {
		m.ShowMsg(f.Path().Current() + " Folder is empty")
		f.Path().Pop()
	}
	m.Info("Path: " + f.Path().Current())
	return nil
}

