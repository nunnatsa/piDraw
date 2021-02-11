package canvas

import "github.com/nunnatsa/piDraw/datatype"

// Cursor is the location of the current pixel
type Cursor struct {
	X     uint8          `json:"x"`
	Y     uint8          `json:"y"`
	Color datatype.Color `json:"color"`
}

// MoveUp moves the curser to the pixel above
func (c *Cursor) MoveUp() {
	if c.Y > 0 {
		c.Y--
	}
}

// MoveDown moves the curser to the pixel below
func (c *Cursor) MoveDown() {
	if c.Y < canvasHight-1 {
		c.Y++
	}
}

// MoveLeft moves the curser to the pixel left
func (c *Cursor) MoveLeft() {
	if c.X > 0 {
		c.X--
	}
}

// MoveRight moves the curser to the pixel right
func (c *Cursor) MoveRight() {
	if c.X < canvasWidth-1 {
		c.X++
	}
}

// SetColor set the working Color. This is the Color that will used to paint a pixel in the Canvas.
func (c *Cursor) SetColor(clr datatype.Color) {
	c.Color = clr
}

// GetColor returns the cursor Color
func (c Cursor) GetColor() datatype.Color {
	return c.Color
}
