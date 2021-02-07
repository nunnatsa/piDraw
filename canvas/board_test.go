package canvas

import "testing"

func TestNewBoard(t *testing.T) {
	b := NewBoard(nil, nil)

	if b.Window.matrix[4][4] != 0 {
		t.Errorf("board should be initialized, but b.Window[4][4] is %d", b.Window.matrix[4][4])
	}

	b.Cursor.SetColor(15)
	err := b.Canvas.Set(b.Cursor)
	if err != nil {
		t.Errorf("should not return error; error is %v", err)
	}

	if b.Window.matrix[4][4] != 15 {
		t.Errorf("b.Window[4][4] should be 15, but it's %d", b.Window.matrix[4][4])
	}
}

func TestMoveUpWindowMidleCursorMidle(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: centerX, Y: centerY, Color: 45},
		Window: c.prepareWindow(windowSize, windowSize),
	}

	b.MoveUp()
	if b.Cursor.Y != centerY-1 {
		t.Error("cursor should move up")
	}

	if b.Window.Y != windowSize {
		t.Error("window should not move up")
	}
}

func TestMoveUpWindowMidleCursorTop(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: centerX, Y: centerY, Color: 45},
		Window: c.prepareWindow(windowSize, centerY),
	}

	b.MoveUp()
	if b.Cursor.Y != centerY-1 {
		t.Error("cursor should move up")
	}

	if b.Window.Y != centerY-1 {
		t.Error("window should move up")
	}
}

func TestMoveUpWindowTopCursorMidle(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: centerX, Y: 3, Color: 45},
		Window: c.prepareWindow(centerX, 0),
	}

	b.MoveUp()
	if b.Cursor.Y != 2 {
		t.Error("cursor should move up")
	}

	if b.Window.Y != 0 {
		t.Error("window should not move up")
	}
}

func TestMoveUpWindowTopCursorTop(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: centerX, Y: 0, Color: 45},
		Window: c.prepareWindow(windowSize, 0),
	}

	b.MoveUp()
	if b.Cursor.Y != 0 {
		t.Error("cursor should not move up")
	}

	if b.Window.Y != 0 {
		t.Error("window should not move up")
	}
}

func TestMoveDownWindowMidleCursorMidle(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: centerX, Y: centerY, Color: 45},
		Window: c.prepareWindow(windowSize, windowSize),
	}

	b.MoveDown()
	if b.Cursor.Y != centerY+1 {
		t.Error("cursor should move down")
	}

	if b.Window.Y != windowSize {
		t.Error("window should not move down")
	}
}

func TestMoveDownWindowMidleCursorBottom(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{
			X:     centerX,
			Y:     windowSize + windowSize - 1,
			Color: 45,
		},
		Window: c.prepareWindow(windowSize, windowSize),
	}

	b.MoveDown()
	if b.Cursor.Y != windowSize+windowSize {
		t.Errorf("cursor should move down; b.Cursor.Y = %d", b.Cursor.Y)
	}

	if b.Window.Y != windowSize+1 {
		t.Errorf("window should move down; b.Window.Y = %d", b.Window.Y)
	}
}

func TestMoveDownWindowBottomCursorMidle(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{
			X:     centerX,
			Y:     canvasHight - 5,
			Color: 45,
		},
		Window: c.prepareWindow(centerX, canvasHight-windowSize),
	}

	b.MoveDown()
	if b.Cursor.Y != canvasHight-4 {
		t.Error("cursor should move down")
	}

	if b.Window.Y != canvasHight-windowSize {
		t.Error("window should not move down")
	}
}

func TestMoveDownWindowTopCursorBottom(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: centerX, Y: canvasHight - 1, Color: 45},
		Window: c.prepareWindow(windowSize, canvasHight-windowSize),
	}

	b.MoveDown()
	if b.Cursor.Y != canvasHight-1 {
		t.Error("cursor should not move down")
	}

	if b.Window.Y != canvasHight-windowSize {
		t.Error("window should not move down")
	}
}

func TestMoveLeftWindowMidleCursorMidle(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: centerX, Y: centerY, Color: 45},
		Window: c.prepareWindow(windowSize, windowSize),
	}

	b.MoveLeft()
	if b.Cursor.X != centerX-1 {
		t.Error("cursor should move left")
	}

	if b.Window.X != windowSize {
		t.Error("window should not move left")
	}
}

func TestMoveLeftWindowMidleCursorLeft(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: windowSize, Y: centerY, Color: 45},
		Window: c.prepareWindow(windowSize, windowSize),
	}

	b.MoveLeft()
	if b.Cursor.X != windowSize-1 {
		t.Error("cursor should move left")
	}

	if b.Window.X != windowSize-1 {
		t.Error("window should move left")
	}
}

func TestMoveLeftWindowLeftCursorMidle(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: 3, Y: centerY, Color: 45},
		Window: c.prepareWindow(0, centerY),
	}

	b.MoveLeft()
	if b.Cursor.X != 2 {
		t.Error("cursor should move left")
	}

	if b.Window.X != 0 {
		t.Error("window should not move left")
	}
}

func TestMoveLeftWindowLeftCursorLeft(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: 0, Y: centerY, Color: 45},
		Window: c.prepareWindow(0, centerY),
	}

	b.MoveLeft()
	if b.Cursor.X != 0 {
		t.Error("cursor should not move left")
	}

	if b.Window.X != 0 {
		t.Error("window should not move left")
	}
}

func TestMoveRightWindowMidleCursorMidle(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: centerX, Y: centerY, Color: 45},
		Window: c.prepareWindow(windowSize, windowSize),
	}

	b.MoveRight()
	if b.Cursor.X != centerX+1 {
		t.Error("cursor should move right")
	}

	if b.Window.X != windowSize {
		t.Error("window should not move right")
	}
}

func TestMoveRightWindowMidleCursorRight(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{
			X:     windowSize + windowSize - 1,
			Y:     centerY,
			Color: 45,
		},
		Window: c.prepareWindow(windowSize, windowSize),
	}

	b.MoveRight()
	if b.Cursor.X != windowSize+windowSize {
		t.Errorf("cursor should move right; b.Cursor.X = %d", b.Cursor.X)
	}

	if b.Window.X != windowSize+1 {
		t.Errorf("window should move right; b.Window.X = %d", b.Window.X)
	}
}

func TestMoveRightWindowRightCursorMidle(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{
			X:     canvasWidth - 5,
			Y:     centerY,
			Color: 45,
		},
		Window: c.prepareWindow(canvasHight-windowSize, centerY),
	}

	b.MoveRight()
	if b.Cursor.X != canvasWidth-4 {
		t.Error("cursor should move right")
	}

	if b.Window.X != canvasWidth-windowSize {
		t.Error("window should not move right")
	}
}

func TestMoveRightWindowRightCursorRight(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: canvasWidth - 1, Y: centerY, Color: 45},
		Window: c.prepareWindow(canvasWidth-windowSize, centerY),
	}

	b.MoveRight()
	if b.Cursor.X != canvasWidth-1 {
		t.Error("cursor should not move right")
	}

	if b.Window.X != canvasWidth-windowSize {
		t.Error("window should not move right")
	}
}

func TestBoardSet(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: windowSize + 1, Y: windowSize - 1, Color: 45},
		Window: c.prepareWindow(windowSize, windowSize),
	}

	b.DrawPixel()

	if (*b.Canvas)[windowSize-1][windowSize+1] != 45 {
		t.Errorf("should set (%d, %d) to 45, but it's %d instead", windowSize+1, windowSize-1, (*b.Canvas)[windowSize-1][windowSize+1])
	}
}

func TestBoardDelete(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: windowSize + 1, Y: windowSize - 1, Color: 45},
		Window: c.prepareWindow(windowSize, windowSize),
	}

	b.DeletePixel()

	if (*b.Canvas)[windowSize-1][windowSize+1] != 0 {
		t.Errorf("should Delete (%d, %d), but it's %d instead", windowSize+1, windowSize-1, (*b.Canvas)[windowSize-1][windowSize+1])
	}
}

func TestResetBoard(t *testing.T) {
	c := getTestCanvas()
	b := &Board{
		Canvas: c,
		Cursor: &Cursor{X: 0, Y: 0, Color: 45},
		Window: c.prepareWindow(0, 0),
	}
	b.Reset()

	for y := 0; y < canvasHight; y++ {
		for x := 0; x < canvasWidth; x++ {
			if (*b.Canvas)[y][x] != 0 {
				t.Errorf("(%d, %d) shold be 0 but it %d", x, y, (*b.Canvas)[y][x])
			}
		}
	}

	if b.Cursor.X != centerX {
		t.Errorf("b.Cursor.X should be %d but it's %d", centerX, b.Cursor.X)
	}

	if b.Cursor.Y != centerY {
		t.Errorf("b.Cursor.Y should be %d but it's %d", centerY, b.Cursor.Y)
	}

	if b.Cursor.Color != 0xFFFFFF {
		t.Errorf("b.Cursor.Color should be %d but it's %d", 0xFFFFFF, b.Cursor.Color)
	}

	if b.Window.X != windowSize {
		t.Errorf("b.Window.X should be %d but it's %d", windowSize, b.Window.X)
	}

	if b.Window.Y != windowSize {
		t.Errorf("b.Window.Y should be %d but it's %d", windowSize, b.Window.Y)
	}
}
