package canvas

import (
	"testing"
)

func TestSet(t *testing.T) {
	c := getTestCanvas()

	err := c.Set(&Cursor{1, 2, 3})
	if err != nil {
		t.Error("Should not return error")
	}

	if c[2][1] != 3 {
		t.Error("Pixel should be 3")
	}

	err = c.Set(&Cursor{canvasWidth, 2, 3})
	if err == nil {
		t.Error("Should return error")
	}

	err = c.Set(&Cursor{1, canvasHeight, 3})
	if err == nil {
		t.Error("Should return error")
	}

	err = c.Set(&Cursor{canvasWidth + 1, 2, 3})
	if err == nil {
		t.Error("Should return error")
	}

	err = c.Set(&Cursor{1, canvasHeight + 1, 1})
	if err == nil {
		t.Error("Should return error")
	}

	err = c.Set(&Cursor{canvasWidth - 1, 1, 1})
	if err != nil {
		t.Error("Should not return error")
	}

	if c[1][canvasWidth-1] != 1 {
		t.Error("Pixel should be 1")
	}

	err = c.Set(&Cursor{1, canvasHeight - 1, 1})
	if err != nil {
		t.Error("Should not return error")
	}

	if c[canvasHeight-1][1] != 1 {
		t.Error("Pixel should be 1")
	}
}

func TestDelete(t *testing.T) {
	c := getTestCanvas()

	err := c.Delete(&Cursor{1, 2, 3})
	if err != nil {
		t.Error("Should not return error")
	}

	if c[2][1] != 0 {
		t.Error("Pixel should be 0")
	}

	err = c.Delete(&Cursor{canvasWidth, 2, 3})
	if err == nil {
		t.Error("Should return error")
	}

	err = c.Delete(&Cursor{1, canvasHeight, 3})
	if err == nil {
		t.Error("Should return error")
	}

	err = c.Delete(&Cursor{canvasWidth + 1, 2, 3})
	if err == nil {
		t.Error("Should return error")
	}

	err = c.Delete(&Cursor{1, canvasHeight + 1, 1})
	if err == nil {
		t.Error("Should return error")
	}

	err = c.Delete(&Cursor{canvasWidth - 1, 1, 1})
	if err != nil {
		t.Error("Should not return error")
	}

	if c[1][canvasWidth-1] != 0 {
		t.Error("Pixel should be 0")
	}

	err = c.Delete(&Cursor{1, canvasHeight - 1, 1})
	if err != nil {
		t.Error("Should not return error")
	}

	if c[canvasHeight-1][1] != 0 {
		t.Error("Pixel should be 0")
	}
}

func TestReset(t *testing.T) {
	c := getTestCanvas()

	c.Reset()

	for y := 0; uint8(y) < canvasHeight; y++ {
		for x := 0; uint8(x) < canvasWidth; x++ {
			if c[y][x] != 0 {
				t.Errorf("(%d, %d) shold be 0 but it %d", x, y, c[y][x])
			}
		}
	}
}

func TestPrepareWindow(t *testing.T) {
	c := getTestCanvas()

	w := c.prepareWindow(4, 6)

	if w.matrix[1][2] != 7 {
		t.Errorf("w[1][2] should be 7, but it's %d", w.matrix[1][2])
	}
}

func getTestCanvas() Canvas {
	return Canvas{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
	}
}
