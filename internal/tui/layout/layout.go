package layout

import "github.com/rivo/tview"

type Layout interface {
	SetContent(tview.Primitive)
	SetSidebar(tview.Primitive)
	GetLayout() *tview.Grid
	SetFocused(tview.Primitive)
	SetStatus(string)
	GetFocused() *tview.Primitive
}

type position struct {
	row    int
	col    int
	width  int
	height int
}

type layout struct {
	grid    *tview.Grid
	content tview.Primitive
	sidebar tview.Primitive
	focused *tview.Primitive
	status  *tview.TextView
}

var (
	statusPos  position = position{row: 1, col: 0, width: 2, height: 1}
	contentPos position = position{row: 0, col: 0, width: 1, height: 1}
	sidePos    position = position{row: 0, col: 1, width: 1, height: 1}
)

func New() Layout {
	grid := tview.NewGrid().
		SetRows(-1, 3).
		SetColumns(-3, -1)

	status := tview.NewTextView()
	status.SetBorder(true).SetTitle("Status bar")

	grid.
		AddItem(status, statusPos.row, statusPos.col, statusPos.height, statusPos.width, 0, 0, false)
		// AddItem(tview.NewBox().
		// SetBorder(true).
		// SetTitle("Side bar"), sidePos.row, sidePos.col, sidePos.height, sidePos.width, 0, 0, false)

	return &layout{
		grid:   grid,
		status: status,
	}
}

func (l *layout) SetFocused(primitive tview.Primitive) {
	l.focused = &primitive
}

func (l layout) GetFocused() *tview.Primitive {
	return l.focused
}

func (l layout) GetLayout() *tview.Grid {
	return l.grid
}

func (l *layout) SetContent(content tview.Primitive) {
	l.grid.RemoveItem(l.content)
	l.grid.AddItem(content, contentPos.row, contentPos.col, contentPos.height, contentPos.width, 0, 0, false)
	l.content = content
}

func (l *layout) SetStatus(msg string) {
	// l.grid.RemoveItem(l.sidebar)
	// l.grid.AddItem(sidebar, sidePos.row, sidePos.col, sidePos.height, sidePos.width, 0, 0, false)
	// l.sidebar = sidebar
	l.status.SetText(msg)
}

func (l *layout) SetSidebar(sidebar tview.Primitive) {
	l.grid.RemoveItem(l.sidebar)
	l.grid.AddItem(sidebar, sidePos.row, sidePos.col, sidePos.height, sidePos.width, 0, 0, false)
	l.sidebar = sidebar
}
