package main

import (
	"./filemanager"
	"./menu"
)

func enter(mn *menu.Menu) error {
	_, activePath := mn.GetActive()
	if filemanager.IsDir(activePath) {
		fm.Path().Set(activePath)
		files, err := fm.Ls()
		if err != nil {
			return err
		}

		err = refresh(mn, &fm, files)
		if err != nil {
			fm.Path().Pop()
			mn.ShowMsg(err.Error())
		}
	} else {
		// c := make(chan string)
		// worker := &Worker{Command: "xdg-open", Args: fm.Path().Current(), Output: c}
		// go worker.Run()
		mn.ShowMsg(activePath + " is file")
	}
	return nil
}
func back(mn *menu.Menu) error {
	curr := fm.Path().Current()
	fm.Path().Pop()
	files, err := fm.Ls()

	if err != nil {
		return err
	}

	err = refresh(mn, &fm, files)
	if err != nil {
		fm.Path().Set(curr)
		mn.ShowMsg(curr + " is the root directory")
	}
	return nil

}

func exit(m *menu.Menu) error {
	menu.Close()
	return nil
}

func setBookmark(m *menu.Menu) error {
	title, path := m.GetActive()
	bookmark.Add(title, path)
	m.ShowMsg("Bookmark added. Press Ctrl+/ to see all bookmarks")
	return nil
}

func showBookmarks(m *menu.Menu) error {
	bookmarks, err := bookmark.Get()
	if err != nil {
		m.ShowMsg(err.Error())
	}
	files := []filemanager.File{}
	for t, p := range bookmarks {
		files = append(files, filemanager.File{
			Name: t,
			Path: p,
		})
	}
	refresh(m, &fm, files)
	return nil
}
