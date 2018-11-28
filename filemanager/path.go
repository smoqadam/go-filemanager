package filemanager

import (
	"path/filepath"
	"strings"
)

type Path struct {
	full []string
	curr string
}

func (p *Path) Push(path string) []string {
	p.full = append(p.full, path)
	return p.full
}

func (p *Path) Pop() string {
	if len(p.full) <= 0 {
		return "/"
	}
	x, a := p.full[len(p.full)-1], p.full[:len(p.full)-1]
	p.full = a
	return x
}

func (p *Path) Current() string {
	//TODO: Fix this shit
	return filepath.Clean("/" + strings.Join(p.full, "/"))
}

func (p *Path) Set(path string) {
	p.full = strings.Split(path, "/")
}
