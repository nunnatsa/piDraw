package datatype

import "fmt"

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

func NewDisplayMessage(mat [][]Color, x, y uint8) *DisplayMessage {
	return &DisplayMessage{
		Screen:  mat,
		CursorX: x,
		CursorY: y,
	}
}

func (dm DisplayMessage) String() string {
	return fmt.Sprintf("screen: %v\ncursor: (%d, %d)", dm.Screen, dm.CursorX, dm.CursorY)
}

// Color is the Color of one pixel in the Canvas
type Color uint32

type ClientEventType int32

const (
	EventTypeReset ClientEventType = iota
	EventTypeColorChange
	EventClientRegistered
)

type ClientEvent struct {
	EventType ClientEventType
	Data      interface{}
}
