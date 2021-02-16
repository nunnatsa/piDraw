package controller

import (
	"github.com/nunnatsa/piDraw/datatype"
	"github.com/nunnatsa/piDraw/hat"
	"github.com/nunnatsa/piDraw/notifier"
	"testing"
)

const (
	windowSize = uint8(8)
)

func TestDo(t *testing.T) {
	myHat := hat.NewMock()
	defer myHat.Close()
	c := NewController(3, 3, notifier.NewNotifier(), myHat)

	c.Start()
	msg := myHat.GetLastMsg()
	if msg.CursorY != windowSize/2 {
		t.Errorf("msg.CursorY should be %d but it's %d", windowSize/2, msg.CursorY)
	}
	if c.board.Cursor.Y != windowSize+windowSize/2 {
		t.Errorf("c.boardCursor.Y should be %d but it's %d", windowSize+windowSize/2, c.board.Cursor.Y)
	}

	myHat.GoDown()
	msg = myHat.GetLastMsg()
	if msg.CursorY != 5 {
		t.Errorf("msg.CursorY should be 5 but it's %d", msg.CursorY)
	}
	if c.board.Cursor.Y != windowSize+5 {
		t.Errorf("c.boardCursor.Y should be %d but it's %d", windowSize+5, c.board.Cursor.Y)
	}

	myHat.GoUp()
	msg = myHat.GetLastMsg()
	if msg.CursorY != 4 {
		t.Errorf("msg.CursorY should be 4 but it's %d", msg.CursorY)
	}
	if c.board.Cursor.Y != windowSize+4 {
		t.Errorf("c.boardCursor.Y should be %d but it's %d", windowSize+4, c.board.Cursor.Y)
	}

	myHat.GoRight()
	msg = myHat.GetLastMsg()
	if msg.CursorX != 5 {
		t.Errorf("msg.CursorX should be 5 but it's %d", msg.CursorX)
	}
	if c.board.Cursor.X != windowSize+5 {
		t.Errorf("c.boardCursor.X should be %d but it's %d", windowSize+5, c.board.Cursor.X)
	}

	myHat.GoLeft()
	msg = myHat.GetLastMsg()
	if msg.CursorX != 4 {
		t.Errorf("msg.CursorX should be 4 but it's %d", msg.CursorX)
	}
	if c.board.Cursor.X != windowSize+4 {
		t.Errorf("c.board.Cursor.X should be %d but it's %d", windowSize+4, c.board.Cursor.X)
	}

	// make sure the original value is there, before changing
	if c.board.Canvas[windowSize+4][windowSize+4] != 0 {
		t.Errorf("Canvas[%d][%d] should be 0 but it's %d", windowSize+4, windowSize+4, c.board.Canvas[windowSize+4][windowSize+4])
	}
	c.board.Cursor.Color = 123456
	myHat.Press()
	msg = myHat.GetLastMsg()
	if msg.CursorX != 4 {
		t.Errorf("msg.CursorX should be 4 but it's %d", msg.CursorX)
	}
	if msg.CursorY != 4 {
		t.Errorf("msg.CursorY should be 4 but it's %d", msg.CursorY)
	}
	if msg.Screen[4][4] != 123456 {
		t.Errorf("msg.Screen[4][4] should be 123456 but it's %d", msg.Screen[4][4])
	}
	if c.board.Cursor.X != windowSize+4 {
		t.Errorf("c.board.Cursor.X should be %d but it's %d", windowSize+4, c.board.Cursor.X)
	}
	if c.board.Cursor.Y != windowSize+4 {
		t.Errorf("c.board.Cursor.Y should be %d but it's %d", windowSize+4, c.board.Cursor.Y)
	}
	if c.board.Canvas[windowSize+4][windowSize+4] != 123456 {
		t.Errorf("Canvas[%d][%d] should be 123456 but it's %d", windowSize+4, windowSize+4, c.board.Canvas[windowSize+4][windowSize+4])
	}

	c.clientEvents <- datatype.ClientEvent{EventType: datatype.ClientEventType(0xFFFF)}
	if len(c.screenEvents) > 0 {
		t.Errorf("Should do nothing, but produced screen event %v", <-c.screenEvents)
	}

	c.clientEvents <- datatype.ClientEvent{EventType: datatype.EventTypeReset}
	msg = myHat.GetLastMsg()
	if msg.CursorX != 4 {
		t.Errorf("msg.CursorX should be 4 but it's %d", msg.CursorX)
	}
	if msg.CursorY != 4 {
		t.Errorf("msg.CursorY should be 4 but it's %d", msg.CursorY)
	}
	if msg.Screen[4][4] != 0 {
		t.Errorf("msg.Screen[4][4] should be 0 but it's %d", msg.Screen[4][4])
	}
	if c.board.Cursor.X != windowSize+4 {
		t.Errorf("c.board.Cursor.X should be %d but it's %d", windowSize+4, c.board.Cursor.X)
	}
	if c.board.Cursor.Y != windowSize+4 {
		t.Errorf("c.board.Cursor.Y should be %d but it's %d", windowSize+4, c.board.Cursor.Y)
	}
	if c.board.Canvas[windowSize+4][windowSize+4] != 0 {
		t.Errorf("Canvas[%d][%d] should be 0 but it's %d", windowSize+4, windowSize+4, c.board.Canvas[windowSize+4][windowSize+4])
	}
}
