package canvas

import (
	"fmt"
	"github.com/nunnatsa/piDraw/datatype"
)

const (
	canvasHight = windowSize * 3
	canvasWidth = windowSize * 3
	centerX     = canvasWidth / 2
	centerY     = canvasHight / 2
)

// Canvas is the virtual Canvas for the drawing
type Canvas [][]datatype.Color

func newCanvas() *Canvas {
	c := make([][]datatype.Color, canvasHight)
	for y := 0; y < canvasHight; y++ {
		c[y] = make([]datatype.Color, canvasWidth)
	}

	return ((*Canvas)(&c))
}

func (c Canvas) prepareWindow(x, y uint8) *Window {
	m := matrix(make([][]datatype.Color, windowSize))

	for i := uint8(0); i < windowSize; i++ {
		m[i] = c[y+i][x : x+windowSize]
	}

	return &Window{matrix: m, X: x, Y: y}
}

// Set set the Color of one pixel in the Canvas
func (c *Canvas) Set(cr *Cursor) error {
	if cr.X >= canvasHight || cr.Y >= canvasWidth {
		return fmt.Errorf(`(%d, %d) is out of the Canvas size`, cr.X, cr.Y)
	}

	(*c)[cr.Y][cr.X] = cr.Color

	return nil
}

// Delete deletes one pixel in the Canvas
func (c *Canvas) Delete(cr *Cursor) error {
	if cr.X >= canvasHight || cr.Y >= canvasWidth {
		return fmt.Errorf(`(%d, %d) is out of the Canvas size`, cr.X, cr.Y)
	}

	(*c)[cr.Y][cr.X] = 0

	return nil
}

// Reset set all the pixel in the Canvas to zero
func (c *Canvas) Reset() {
	for y := 0; y < canvasHight; y++ {
		for x := 0; x < canvasWidth; x++ {
			(*c)[y][x] = 0
		}
	}
}
