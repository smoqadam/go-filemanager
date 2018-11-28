package menu

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

type Window struct {
	width        int
	height       int
	startListCol int
	startListRow int
	infoCol      int
	infoRow      int
	errCol       int
	errRow       int
}

type Menu struct {
	items            []MenuItem
	itemsActiveIndex int
	loadedItems      []MenuItem
	activeIndex      int
	eventKey         map[termbox.Key]func(*Menu) error
	window           *Window
	info             string
}

func New(mItems []MenuItem) Menu {

	// var mItems []MenuItem
	// for _, mi := range mItem {
	// 	mItems = append(mItems, MenuItem{
	// 		Title: mi,
	// 		Value:
	// 	})
	// }

	m := Menu{
		items:            mItems,
		eventKey:         make(map[termbox.Key]func(*Menu) error),
		activeIndex:      0,
		itemsActiveIndex: 0,
	}

	return m
}

func (m *Menu) Render() {
	err := termbox.Init()

	width, height := termbox.Size()
	m.window = &Window{
		width:        width,
		height:       height - 5,
		startListCol: 2,
		startListRow: 4,
		infoCol:      2,
		infoRow:      4,
		errCol:       height - 2,
		errRow:       4,
	}
	// window.height = window.height - 5
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	lItems := m.items
	if len(m.items) > m.window.height {
		lItems = m.items[0:m.window.height]
	}
	m.loadedItems = lItems

	m.setEvents()
	m.loop()
}

func (m *Menu) setEvents() {
	m.AddEvent(termbox.KeyCtrlC, func(m *Menu) error {
		Close()
		return nil
	})

	m.AddEvent(termbox.KeyArrowDown, func(m *Menu) error {
		m.activeIndex++
		m.itemsActiveIndex++

		if m.activeIndex >= len(m.loadedItems) {
			if len(m.items) > len(m.loadedItems) {
				if m.itemsActiveIndex >= len(m.items) {
					m.itemsActiveIndex = 0
					m.loadedItems = m.items[m.itemsActiveIndex:m.window.height]
					m.activeIndex = 0
				} else {
					m.loadedItems = m.loadedItems[1:]
					m.loadedItems = append(m.loadedItems, m.items[m.itemsActiveIndex])
					m.activeIndex = m.window.height - 1
				}
			} else {
				m.activeIndex = 0
			}
		}
		return nil

	})

	m.AddEvent(termbox.KeyArrowUp, func(m *Menu) error {
		m.activeIndex--
		m.itemsActiveIndex--

		if m.activeIndex <= 0 {
			if len(m.items) > len(m.loadedItems) {
				if m.itemsActiveIndex < 0 {
					m.itemsActiveIndex = len(m.items) - 1
					m.loadedItems = m.items[m.itemsActiveIndex-m.window.height : m.itemsActiveIndex]
					m.activeIndex = len(m.loadedItems) - 1
				} else {
					a := m.items[m.itemsActiveIndex]
					x := []MenuItem{a}
					x = append(x, m.loadedItems[:len(m.loadedItems)-1]...)
					m.loadedItems = append([]MenuItem{a}, m.loadedItems[:len(m.loadedItems)-1]...)
					m.activeIndex = 0
				}
			} else {
				m.activeIndex = len(m.loadedItems) - 1
			}
		}
		return nil
	})

}

func (m *Menu) loop() {
	event := make(chan termbox.Event)
	go func() {
		for {
			event <- termbox.PollEvent()
		}
	}()
	for {
		select {
		case ev := <-event:
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

			if fn, ok := m.eventKey[ev.Key]; ok {
				fn(m)
			}

			m.Print()
			termbox.Flush()
		default:
			m.Print()
			termbox.Flush()
		}
	}
}

func (m *Menu) Print() {
	for i, mi := range m.loadedItems {

		if m.activeIndex == i {
			mi.bg = termbox.ColorBlue
		}
		mi.Print(m.window.startListRow, i+m.window.startListCol)
	}
	m.Header(m.info)
}

func (m *Menu) AddEvent(e termbox.Key, fn func(*Menu) error) {
	m.eventKey[e] = fn
}

func (m *Menu) AddItem(mi MenuItem) {
	m.items = append(m.items, mi)
}

func (m *Menu) GetActive() (string, string) {
	if m.activeIndex <= 0 {
		m.activeIndex = 0
	}
	if m.activeIndex >= len(m.items) {
		m.activeIndex = len(m.items)
	}
	if len(m.loadedItems) < m.activeIndex {
		return m.loadedItems[0].Title, m.loadedItems[m.activeIndex].Value
	}

	return m.loadedItems[m.activeIndex].Title, m.loadedItems[m.activeIndex].Value
}

func (m *Menu) SetStringItems(items []string) error {

	if len(items) <= 0 {
		return fmt.Errorf("Empty")
	}
	m.items = nil
	m.itemsActiveIndex = 0
	m.activeIndex = 0
	for _, i := range items {
		m.AddItem(MenuItem{
			Title: i,
		})
	}
	lItems := m.items
	if len(m.items) > m.window.height {
		lItems = m.items[0:m.window.height]
	}
	m.loadedItems = lItems
	return nil
}

func (m *Menu) SetItems(items []MenuItem) error {

	if len(items) <= 0 {
		return fmt.Errorf("Empty")
	}
	m.items = nil
	m.itemsActiveIndex = 0
	m.activeIndex = 0
	// for _, i := range items {
	// 	m.AddItem(MenuItem{
	// 		Title: i,
	// 	})
	// }
	m.items = items
	lItems := m.items
	if len(m.items) > m.window.height {
		lItems = m.items[0:m.window.height]
	}
	m.loadedItems = lItems
	return nil
}
func (m *Menu) ShowMsg(msg string) {

	for i, ch := range msg {
		termbox.SetCell(m.window.errRow+i, m.window.errCol, ch, termbox.ColorCyan, termbox.ColorDefault)
	}

}
func (m *Menu) Info(i string) {
	m.info = i
}
func (m *Menu) Header(msg string) {
	for i, ch := range msg {
		termbox.SetCell(m.window.infoRow+i, 0, ch, termbox.ColorCyan, termbox.ColorDefault)
	}
}
func Close() {
	termbox.Close()
	os.Exit(0)

}
