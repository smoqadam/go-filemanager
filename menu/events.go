package menu

func close(m *Menu) {
	Close()
	// return nil
}

func goDown(m *Menu) {
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
	// return nil

}

func goUp(m *Menu) {
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
	// return nil
}
