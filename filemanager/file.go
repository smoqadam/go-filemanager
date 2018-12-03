package filemanager

type File struct {
	Name   string
	Path   string
	Active bool
	IsDir  bool
	Size   int32
}
