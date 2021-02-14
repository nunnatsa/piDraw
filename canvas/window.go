package canvas

import "github.com/nunnatsa/piDraw/datatype"

const (
	windowSize = 8 // hat display is 8X8
)

type Matrix [][]datatype.Color

// Window is an 8X8 Matrix that is a subset of the Canvas, and the current disply in the HAT
type Window struct {
	Matrix `json:"-"`
	// top-left pixel location in the Canvas
	X uint8 `json:"x"`
	Y uint8 `json:"y"`
}
