package canvas

import "testing"

func TestMoveUpMiddle(t *testing.T) {
	c := Cursor{X: 8, Y: 8}
	c.MoveUp()

	if c.X != 8 {
		t.Errorf("X should not be changed, but it's %d", c.X)
	}

	if c.Y != 7 {
		t.Errorf("Y should not decreased by 1, but it's %d", c.Y)
	}
}

func TestMoveUpTop(t *testing.T) {
	c := Cursor{X: 8, Y: 0}
	c.MoveUp()

	if c.X != 8 {
		t.Errorf("X should not be changed, but it's %d", c.X)
	}

	if c.Y != 0 {
		t.Errorf("Y should not be changed, but it's %d", c.Y)
	}
}

func TestMoveUpBottom(t *testing.T) {
	c := Cursor{X: 8, Y: canvasHight - 1}
	c.MoveUp()

	if c.X != 8 {
		t.Errorf("X should not be changed, but it's %d", c.X)
	}

	if c.Y != canvasHight-2 {
		t.Errorf("Y should not be decreased by 1, but it's %d", c.Y)
	}
}

func TestMoveDownMiddle(t *testing.T) {
	c := Cursor{X: 8, Y: 8}
	c.MoveDown()

	if c.X != 8 {
		t.Errorf("X should not be changed, but it's %d", c.X)
	}

	if c.Y != 9 {
		t.Errorf("Y should increased by 1, but it's %d", c.Y)
	}
}

func TestMoveDownTop(t *testing.T) {
	c := Cursor{X: 8, Y: 0}
	c.MoveDown()

	if c.X != 8 {
		t.Errorf("X should not be changed, but it's %d", c.X)
	}

	if c.Y != 1 {
		t.Errorf("Y should increased by 1, but it's %d", c.Y)
	}
}

func TestMoveDownBottom(t *testing.T) {
	c := Cursor{X: 8, Y: canvasHight - 1}
	c.MoveDown()

	if c.X != 8 {
		t.Errorf("X should not be changed, but it's %d", c.X)
	}

	if c.Y != canvasHight-1 {
		t.Errorf("Y should not be dechanged, but it's %d", c.Y)
	}
}

func TestMoveLeftMiddle(t *testing.T) {
	c := Cursor{X: 8, Y: 8}
	c.MoveLeft()

	if c.Y != 8 {
		t.Errorf("Y should not be changed, but it's %d", c.Y)
	}

	if c.X != 7 {
		t.Errorf("X should not decreased by 1, but it's %d", c.X)
	}
}

func TestMoveLeftFromLeft(t *testing.T) {
	c := Cursor{X: 0, Y: 8}
	c.MoveLeft()

	if c.Y != 8 {
		t.Errorf("Y should not be changed, but it's %d", c.Y)
	}

	if c.X != 0 {
		t.Errorf("X should not be changed, but it's %d", c.X)
	}
}

func TestMoveLeftFromRight(t *testing.T) {
	c := Cursor{X: canvasWidth - 1, Y: 8}
	c.MoveLeft()

	if c.Y != 8 {
		t.Errorf("Y should not be changed, but it's %d", c.Y)
	}

	if c.X != canvasWidth-2 {
		t.Errorf("X should not be decreased by 1, but it's %d", c.X)
	}
}

func TestMoveRightMiddle(t *testing.T) {
	c := Cursor{X: 8, Y: 8}
	c.MoveRight()

	if c.Y != 8 {
		t.Errorf("Y should not be changed, but it's %d", c.Y)
	}

	if c.X != 9 {
		t.Errorf("X should increased by 1, but it's %d", c.X)
	}
}

func TestMoveRightFromLeft(t *testing.T) {
	c := Cursor{X: 0, Y: 8}
	c.MoveRight()

	if c.Y != 8 {
		t.Errorf("Y should not be changed, but it's %d", c.Y)
	}

	if c.X != 1 {
		t.Errorf("X should increased by 1, but it's %d", c.X)
	}
}

func TestMoveRightFromRight(t *testing.T) {
	c := Cursor{X: canvasWidth - 1, Y: 8}
	c.MoveRight()

	if c.Y != 8 {
		t.Errorf("Y should not be changed, but it's %d", c.Y)
	}

	if c.X != canvasWidth-1 {
		t.Errorf("X should not be dechanged, but it's %d", c.X)
	}
}

func TestSetColor(t *testing.T) {
	c := Cursor{Color: 3}

	c.SetColor(5)
	if c.Color != 5 {
		t.Errorf("the Color should be 5 but it's %d", c.Color)
	}
}

func TestGetColor(t *testing.T) {
	c := Cursor{Color: 4}
	if clr := c.GetColor(); clr != 4 {
		t.Errorf("the Color should be 4 but it's %d", clr)
	}
}
