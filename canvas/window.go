package canvas

import "github.com/nunnatsa/piDraw/datatype"

const (
	windowSize = 8 // hat display is 8X8
)

type matrix [][]datatype.Color

// Window is an 8X8 matrix that is a subset of the Canvas, and the current disply in the HAT
type Window struct {
	matrix
	// top-left pixel location in the Canvas
	X uint8 `json:"x,omitempty"`
	Y uint8 `json:"y,omitempty"`
}
