package menu

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

func newWindow(width, height int) *Window {
	return &Window{
		width:        width,
		height:       height - 5,
		startListCol: 2,
		startListRow: 4,
		infoCol:      2,
		infoRow:      4,
		errCol:       height - 2,
		errRow:       4,
	}
}

func (w *Window) Width() int {
	return w.width
}

func (w *Window) Height() int {
	return w.height
}
