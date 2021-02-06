package datatype

const (
	Pressed HatEvent = iota
	MoveUp
	MoveLeft
	MoveDown
	MoveRight
)

type HatEvent uint8

type DisplayMessage struct {
	Screen  [][]Color
	CursorX uint8
	CursorY uint8
}

func NewDisplayMessage(mat [][]Color, x, y uint8) *DisplayMessage{
	return &DisplayMessage{
		Screen: mat,
		CursorX: x,
		CursorY: y,
	}
}

// Color is the Color of one pixel in the Canvas
type Color uint32
