package canvas

import (
	"log"
	"sync"
)

type idProvider struct {
	lock    *sync.Mutex
	counter int64
}

func (p *idProvider) getNextID() int64 {
	p.lock.Lock()
	p.counter++
	newID := p.counter
	p.lock.Unlock()
	return newID
}

type Notifier struct {
	clientMap map[int64]chan []byte
	idp       *idProvider
}

func (n *Notifier) Subscribe(ch chan []byte) int64 {
	id := n.idp.getNextID()
	n.clientMap[id] = ch
	log.Println("register new client", id)

	return id
}

func (n *Notifier) Unsubscribe(id int64) {
	log.Println("deregister client", id)
	if ch, ok := n.clientMap[id]; ok {
		close(ch)
		delete(n.clientMap, id)
	}
}

func (n *Notifier) Notify(data []byte) {
	for _, subscriber := range n.clientMap {
		subscriber <- data
	}
}
