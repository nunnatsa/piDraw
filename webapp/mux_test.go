package webapp

import (
	"github.com/gorilla/websocket"
	"github.com/nunnatsa/piDraw/datatype"
	"github.com/nunnatsa/piDraw/notifier"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestColor(t *testing.T) {
	mailbox := notifier.NewNotifier()
	ch := make(chan datatype.ClientEvent)
	ca := NewClientAction(mailbox, 8080, ch)

	body := strings.NewReader(`{"color": 123456}`)
	req := httptest.NewRequest(http.MethodPost, "http://raspberrypi:8080/api/canvas/color", body)
	w := httptest.NewRecorder()
	go ca.setColor(w, req)

	msg := <-ch
	if msg.EventType != datatype.EventTypeColorChange {
		t.Errorf("msg.EventType should be datatype.EventTypeColorChange, but it's %d", msg.EventType)
	}
	color, ok := msg.Data.(datatype.Color)
	if !ok {
		t.Errorf("msg.Data should be with type 'datatype.Color'")
		return
	}
	if color != datatype.Color(123456) {
		t.Errorf("msg.Data should be datatype.Color(123456), but it's %d", color)
	}

}

func TestReset(t *testing.T) {
	mailbox := notifier.NewNotifier()
	ch := make(chan datatype.ClientEvent)
	ca := NewClientAction(mailbox, 8080, ch)

	req := httptest.NewRequest(http.MethodPost, "http://raspberrypi:8080/api/canvas/reset", nil)
	w := httptest.NewRecorder()
	go ca.reset(w, req)

	msg := <-ch
	if msg.EventType != datatype.EventTypeReset {
		t.Errorf("msg.EventType should be datatype.EventTypeReset, but it's %d", msg.EventType)
	}
}

func TestExample(t *testing.T) {
	mailbox := notifier.NewNotifier()
	ch := make(chan datatype.ClientEvent)
	ca := NewClientAction(mailbox, 8080, ch)
	// Create test server with the echo handler.
	s := httptest.NewServer(http.Handler(ca.GetMux()))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	url := u + "/api/canvas/register"
	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	done := make(chan bool)
	msg := <-ch
	if msg.EventType != datatype.EventClientRegistered {
		t.Errorf("msg.EventType should be datatype.EventClientRegistered, but it's %d", msg.EventType)
	}
	go checkMessage(t, ws, "hello test", done)
	mailbox.Notify([]byte("hello test"))
	<-done
}

func checkMessage(t *testing.T, ws *websocket.Conn, expected string, done chan bool) {
	defer close(done)
	_, p, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("%v", err)
	}

	if string(p) != expected {
		t.Errorf("wrong msg: expected = %s; actual = %s", expected, string(p))
	}
}
