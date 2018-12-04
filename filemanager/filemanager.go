package filemanager

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FileManager struct {
	files     []File
	pathStack []string
	path      Path
	sources   []string
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
			Path:   filepath.Clean(f.Path().Current() + "/" + fi.Name()),
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

func (f *FileManager) AddSource(path string) {
	f.sources = append(f.sources, path)
}

func (f *FileManager) Paste(path string) {

	for _, source := range f.sources {
		fileName := strings.Split(source, "/")
		err := copyFileContents(source, path+"/"+fileName[len(fileName)-1])
		if err != nil {
			panic(err)
		}
	}
}

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
