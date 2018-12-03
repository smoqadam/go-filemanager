package menu

import (
	"github.com/nsf/termbox-go"
)

type Section struct {
	Width   int
	Height  int
	X       int
	Y       int
	Title   string
	Footer  string
	Content string
}

func NewSection(width, height, x, y int) Section {
	return Section{
		Title:  "New Sectin",
		Footer: "Section Footer",
		Width:  width,
		Height: height,
		X:      x,
		Y:      y,
	}
}

func (s *Section) SetContent(content string) {
	s.Content = content
}

func (s *Section) Render() {
	var header string
	h := (s.Width + len(s.Title)) / 2
	// t := s.Title[:h]
	for i := 0; i <= s.Width; i++ {
		if i == h {
			header += s.Title
			continue
		}
		header += "-"
	}
	for i, ch := range header {
		termbox.SetCell(s.X+i, s.Y, ch, termbox.ColorDefault, termbox.ColorDefault)
	}

	for i := 1; i < s.Height; i++ {
		termbox.SetCell(s.X, s.Y+i, '|', termbox.ColorDefault, termbox.ColorDefault)
	}

	for i := 1; i < s.Height; i++ {
		termbox.SetCell(s.X+s.Width+len(s.Title), s.Y+i, '|', termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 1; i < s.Width+len(s.Title); i++ {
		termbox.SetCell(s.X+i, s.Y+s.Height, '-', termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(s.X, s.Y+s.Height, '+', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(s.X+len(header), s.Y+s.Height, '+', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(s.X, s.Y, '+', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(s.X+len(header), s.Y, '+', termbox.ColorDefault, termbox.ColorDefault)

	y := 1
	x := s.X + 1
	i := 0
	for _, ch := range s.Content {
		if i >= s.Width && ch == ' ' {
			y++
			// x++
			i = 0
			continue
		}
		if ch == '\n' {
			y++
			i = 0
			continue

		}
		termbox.SetCell(x+i, s.Y+y, ch, termbox.ColorDefault, termbox.ColorDefault)
		i++
	}
}
