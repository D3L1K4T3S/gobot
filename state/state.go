package state

import (
	"image/color"

	"fyne.io/fyne/v2/widget"
)

type Cell int64

const MIN_SIZE = 1
const MAX_SIZE = 10
const COUNT = 5

const (
	Wall Cell = iota
	Resource
	Pit
	Nothing
	Robot
	End
)

var (
	WallColor     = color.NRGBA{A: 0}
	ResourceColor = color.NRGBA{R: 0, G: 255, B: 0, A: 128}
	PitColor      = color.NRGBA{R: 255, G: 0, B: 0, A: 128}
	NothingColor  = color.NRGBA{R: 255, G: 255, B: 255, A: 128}
	RobotColor    = color.NRGBA{R: 0, G: 0, B: 255, A: 128}
	EndColor      = color.NRGBA{R: 228, G: 245, B: 39, A: 128}
)

type AppState struct {
	CellularField []Cell
	Rows          int
	Columns       int
	InputField    []*widget.Select
	score         int
}

func NewAppState() *AppState {
	return &AppState{
		CellularField: make([]Cell, 0),
		Rows:          0,
		Columns:       0,
		InputField:    make([]*widget.Select, 0),
		score:         0,
	}
}
