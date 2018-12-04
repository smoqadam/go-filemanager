package menu

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

const (
	P_CONFIRM string = "CONFIRM"
)

type Menu struct {
	items            []MenuItem
	itemsActiveIndex int
	loadedItems      []MenuItem
	activeIndex      int
	eventKey         map[termbox.Key]func(*Menu)
	window           *Window
	info             string
	currentPage      string
	prevPage         string
	Sections         []*Section
}

func New(mItems []MenuItem) Menu {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	width, height := termbox.Size()
	return Menu{
		items:            mItems,
		eventKey:         make(map[termbox.Key]func(*Menu)),
		activeIndex:      0,
		itemsActiveIndex: 0,
		window:           newWindow(width, height),
	}
}

func (m *Menu) Render() {
	defer termbox.Close()
	mItems := m.items
	if len(m.items) > m.window.height {
		mItems = m.items[0:m.window.height]
	}
	m.loadedItems = mItems

	m.setEvents()
	m.Update()
	m.loop()
}

func (m *Menu) setEvents() {
	m.AddEvent(termbox.KeyCtrlC, close)
	// 	m.AddEvent(termbox.KeyArrowDown, GoDown)
	// 	m.AddEvent(termbox.KeyArrowUp, GoUp)
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
			m.Update()
		}
	}
}

func (m *Menu) Update() {
	for i, mi := range m.loadedItems {
		if m.activeIndex == i {
			mi.bg = termbox.ColorBlue
		}
		mi.Print(m.window.startListRow, i+m.window.startListCol)
	}
	m.Header(m.info)
	for _, s := range m.Sections {
		s.Render()
	}
	termbox.Flush()
}

func (m *Menu) AddEvent(e termbox.Key, fn func(*Menu)) {
	m.eventKey[e] = fn
}

func (m *Menu) AddItem(mi MenuItem) {
	m.items = append(m.items, mi)
}

func (m *Menu) GetActive() MenuItem {
	if m.activeIndex <= 0 {
		m.activeIndex = 0
	}
	if m.activeIndex >= len(m.items) {
		m.activeIndex = len(m.items)
	}
	if len(m.loadedItems) < m.activeIndex {
		return m.loadedItems[0]
	}

	return m.loadedItems[m.activeIndex]
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
func (m *Menu) reset() {
	m.items = nil
	m.itemsActiveIndex = 0
	m.activeIndex = 0

}
func (m *Menu) SetItems(items []MenuItem, currentPage string) error {

	if len(items) <= 0 {
		return fmt.Errorf("Empty")
	}
	m.prevPage = m.currentPage
	m.currentPage = currentPage
	m.reset()
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

func (m *Menu) CurrentPage() string {
	return m.currentPage
}
func (m *Menu) PrevPage() string {
	return m.prevPage

}

func (m *Menu) AddSection(section *Section) {
	m.Sections = append(m.Sections, section)
}

func (m *Menu) Window() *Window {
	return m.window
}
