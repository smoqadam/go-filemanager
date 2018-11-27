package main

import (
	"./filemanager"
	"./menu"
	termbox "github.com/nsf/termbox-go"
)

func main() {
	fm := filemanager.New("/")

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

	m.AddEvent(termbox.KeyEnter, func(mn *menu.Menu) {
		activePath := mn.GetActive()
		fm.Path().Push(activePath)
		if filemanager.IsDir(fm.Path().Current()) {
			err := refresh(mn, &fm, fm.Path().Current())
			if err != nil {
				fm.Path().Pop()
				mn.ShowMsg(err.Error())
			}
		} else {
			mn.ShowMsg(fm.Path().Current() + " is not directory")
			fm.Path().Pop()
		}
	})

	m.AddEvent(termbox.KeyArrowLeft, func(mn *menu.Menu) {
		curr := fm.Path().Current()
		fm.Path().Pop()
		err := refresh(mn, &fm, fm.Path().Current())
		if err != nil {
			fm.Path().Set(curr)
			mn.ShowMsg(curr + " is the first directory")
		}
	})

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

// func List(root string) (error, []File) {
// 	var files []File
// 	fileInfos, err := ioutil.ReadDir(root)
// 	if err != nil {
// 		return err, nil
// 	}
// 	for _, f := range fileInfos {
// 		files = append(files, File{
// 			name:   f.Name(),
// 			path:   root + "/" + f.Name(),
// 			active: false,
// 		})
// 	}

// 	return nil, files
// }
