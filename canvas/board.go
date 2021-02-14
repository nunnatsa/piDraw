package canvas

import (
	"log"
)

var (
	canvasWidth  uint8
	canvasHeight uint8
)

// Board is the draing board.
type Board struct {
	Canvas       Canvas  `json:"canvas,omitempty"`
	Cursor       *Cursor `json:"cursor,omitempty"`
	Window       *Window `json:"window,omitempty"`
	centerX      uint8
	centerY      uint8
}

// NewBoard initiate a new Board
func NewBoard(width, height uint8) *Board {
	canvasWidth = width * windowSize
	canvasHeight = height * windowSize
	centerX := canvasWidth / 2
	centerY := canvasHeight / 2

	c := newCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{
			X:     centerX,
			Y:     centerY,
			Color: 0xFFFFFF,
		},
		Window:  c.prepareWindow(centerX-(windowSize/2), centerY-(windowSize/2)),
		centerX: centerX,
		centerY: centerY,
	}

	return b
}

// MoveUp moves the cursor one pixel up, if not already on top
func (b *Board) MoveUp() {
	b.Cursor.MoveUp()
	if b.Cursor.Y < b.Window.Y {
		b.Window.Y = b.Cursor.Y
		b.Window = b.Canvas.prepareWindow(b.Window.X, b.Window.Y)
	}
}

// MoveDown moves the cursor one pixel down, if not already on bottom
func (b *Board) MoveDown() {
	b.Cursor.MoveDown()
	if b.Cursor.Y > b.Window.Y+windowSize-1 {
		b.Window.Y++
		b.Window = b.Canvas.prepareWindow(b.Window.X, b.Window.Y)
	}
}

// MoveLeft moves the cursor one pixel down, if not already on bottom
func (b *Board) MoveLeft() {
	b.Cursor.MoveLeft()
	if b.Cursor.X < b.Window.X {
		b.Window.X = b.Cursor.X
		b.Window = b.Canvas.prepareWindow(b.Window.X, b.Window.Y)
	}
}

// MoveRight moves the cursor one pixel down, if not already on bottom
func (b *Board) MoveRight() {
	b.Cursor.MoveRight()
	if b.Cursor.X > b.Window.X+windowSize-1 {
		b.Window.X++
		b.Window = b.Canvas.prepareWindow(b.Window.X, b.Window.Y)
	}
}

// DrawPixel writes the pixel the the cursor is currently pointing at, to the Color of the cursor
func (b *Board) DrawPixel() {
	if err := b.Canvas.Set(b.Cursor); err != nil {
		log.Println(err)
	}
}

// DeletePixel writes the pixel the the cursor is currently pointing at, to the Color of the cursor
func (b *Board) DeletePixel() {
	if err := b.Canvas.Delete(b.Cursor); err != nil {
		log.Println(err)
	}
}

// Reset return the board to the initiate state
func (b *Board) Reset() {
	log.Println("Reseting the Canvas")
	b.Canvas = newCanvas()
	b.Cursor = &Cursor{
		X:     b.centerX,
		Y:     b.centerY,
		Color: 0xFFFFFF,
	}
	b.Window = b.Canvas.prepareWindow(b.centerX-(windowSize/2), b.centerY-(windowSize/2))
	log.Println("Canvas is clean now")
}
