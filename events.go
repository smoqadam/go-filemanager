package main

import (
	"net/http"
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

func fileContent(m *menu.Menu) {
	// mItem := m.GetActive()
	// file, _ := ioutil.ReadFile(mItem.Value)

	// contentSection := menu.NewSection(30, 10, m.Window().Width()-45, 7)
	// contentSection.Title = " Content "
	// contentSection.SetContent(string(file))
	// m.AddSection(&contentSection)

	// infoSection.SetContent("Size: " + string(file.Size()))
	// m.AddSection(&infoSection)
	// menu.GoDown(m)

}

func GoUp(m *menu.Menu) {
	menu.GoUp(m)
	mItem := m.GetActive()
	infoSection.SetContent(getFileInfo(mItem.Value))
	m.AddSection(&infoSection)
}

func GoDown(m *menu.Menu) {
	menu.GoDown(m)
	mItem := m.GetActive()
	infoSection.SetContent(getFileInfo(mItem.Value))
	m.AddSection(&infoSection)
}

func getFileInfo(path string) string {
	file, _ := os.Stat(path)
	info := "Size: " + Size(file.Size())
	info += "\n"
	info += "Mode: " + file.Mode().String()
	info += "\n"
	mime, _ := GetFileContentType(path)
	info += "Mime: " + mime

	return info

}

func GetFileContentType(path string) (string, error) {
	// Open File
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err = f.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func copyFile(m *menu.Menu) {
	mItem := m.GetActive()
	fm.AddSource(mItem.Value)
	m.Info(mItem.Value + " added to clipboard. Use Ctrl+p to paste")
}

func pasteFile(m *menu.Menu) {
	src := fm.Path().Current()
	fm.Paste(src)
	m.Info(src + " Copied ")
	files, err := fm.Ls()
	if err != nil {
		m.ShowMsg(err.Error())
	}
	err = refresh(m, &fm, files, PAGE_FILES)
}
