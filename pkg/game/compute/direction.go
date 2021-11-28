package compute

type Direction int

const (
	_ Direction = iota
	Up
	UpperRight
	UpperLeft

	Right
	Left

	Bottom
	BottomRight
	BottomLeft
)

var directions = [...]string{
	Up:          "Up",
	UpperRight:  "UpperRight",
	UpperLeft:   "UpperLeft",
	Right:       "Right",
	Left:        "Left",
	Bottom:      "Bottom",
	BottomRight: "BottomRight",
	BottomLeft:  "BottomLeft",
}

func (d Direction) String() string {
	return directions[d]
}
