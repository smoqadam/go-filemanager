package main

import (
	"os"

	"./filemanager"
	"./menu"
)

var selectedMItem menu.MenuItem

func brows(m *menu.Menu) {
	mItem := m.GetActive()
	if filemanager.IsDir(mItem.Value) {
		fm.Path().Set(mItem.Value)
		files, err := fm.Ls()
		if err != nil {
			// return err
		}

		err = refresh(m, &fm, files, PAGE_FILES)
		if err != nil {
			fm.Path().Pop()
			m.ShowMsg(err.Error())
		}
	} else {
		// TODO: Add file opener
		// c := make(chan string)
		// worker := &Worker{Command: "xdg-open", Args: fm.Path().Current(), Output: c}
		// go worker.Run()
		m.ShowMsg(mItem.Value + " is file")
	}
}
func enter(m *menu.Menu) {

	if m.CurrentPage() == PAGE_CONFIRM {
		mItem := m.GetActive()
		if mItem.Title == "Yes" {
			if m.PrevPage() == PAGE_BOOKMARKS {
				m.ShowMsg("Bookmark " + selectedMItem.Title + " deleted!")
				bookmark.Delete(selectedMItem.Title)
				if _, err := bookmark.Get(); err != nil {
					back(m)
				}

				showBookmarks(m)
				return
			} else {
				m.Info("FILE " + selectedMItem.Value + " deleted!")
				delFile(m, selectedMItem.Value)
				return
			}
		}
	}
	brows(m)

}
func delFile(m *menu.Menu, p string) {
	os.Remove(p)
	curr := fm.Path().Current()
	files, err := fm.Ls()

	if err != nil {
		// return err
	}

	err = refresh(m, &fm, files, PAGE_FILES)
	if err != nil {
		fm.Path().Set(curr)
		m.ShowMsg(curr + " is the root directory")
	}

}

func back(m *menu.Menu) {
	curr := fm.Path().Current()
	fm.Path().Pop()
	files, err := fm.Ls()
	if err != nil {
		// return err
	}

	err = refresh(m, &fm, files, PAGE_FILES)
	if err != nil {
		fm.Path().Set(curr)
		m.ShowMsg(curr + " is the root directory")
	}
}

func exit(m *menu.Menu) {
	menu.Close()
}

func setBookmark(m *menu.Menu) {
	mItem := m.GetActive()
	bookmark.Add(mItem.Title, mItem.Value)
	m.ShowMsg("Bookmark added. Press Ctrl+/ to see all bookmarks")

}

func showBookmarks(m *menu.Menu) {
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
	refresh(m, &fm, files, PAGE_BOOKMARKS)

}

func del(m *menu.Menu) {
	selectedMItem = m.GetActive()
	showConfirm(m, selectedMItem.Value)
}
