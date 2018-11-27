package filemanager

import (
	"io/ioutil"
	"os"
)

type FileManager struct {
	files     []File
	pathStack []string
	path      Path
}

func New(root string) FileManager {
	f := FileManager{}
	f.Path().Push(root)
	files, err := f.Ls()
	if err != nil {
		f.Path().Pop()
		panic(err.Error() + " :: " + f.Path().Current())
	}
	f.files = files
	return f
}

func (f *FileManager) Path() *Path {
	return &f.path
}

func (f *FileManager) Ls() ([]File, error) {
	var files []File
	fileInfos, err := ioutil.ReadDir(f.Path().Current())
	if err != nil {
		return nil, err
	}
	for _, fi := range fileInfos {
		files = append(files, File{
			Name:   fi.Name(),
			Path:   f.Path().Current() + "/" + fi.Name(),
			Active: false,
		})
	}
	return files, nil
}

func (f *FileManager) Cd(p string) {
	f.Path().Set(p)
}

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	mode := fi.Mode()
	return mode.IsDir()
}
