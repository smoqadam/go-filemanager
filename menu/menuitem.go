package menu

import termbox "github.com/nsf/termbox-go"

type MenuItem struct {
	title string
	value string
	bg    termbox.Attribute
	fg    termbox.Attribute
}

func (m *MenuItem) Print(col int, row int) {
	for i, ch := range m.title {
		termbox.SetCell(col+i, row, ch, m.fg, m.bg)
	}

}
