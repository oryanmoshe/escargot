package splash

import (
	"github.com/rivo/tview"
)

type Splash struct {
	title    string
	grid     *tview.Grid
	textview *tview.TextView
}

func New(title string) *Splash {
	grid := tview.NewGrid().SetRows(0).SetColumns(0)
	txtviw := tview.NewTextView() //.SetTextAlign(tview.AlignCenter)
	txtviw.SetBorder(true).SetTitle(title)
	grid.AddItem(txtviw, 0, 0, 1, 1, 0, 0, false)

	return &Splash{
		title:    title,
		grid:     grid,
		textview: txtviw,
	}
}

func (s *Splash) SetContent(text []byte) {
	s.textview.SetText(string(text))
}

func (s Splash) GetGrid() *tview.Grid {
	return s.grid
}
