package notifier

import (
	"log"
	"sync"
	"time"
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
	lock      *sync.Mutex
}

func NewNotifier() *Notifier {
	return &Notifier{
		clientMap: make(map[int64]chan []byte),
		idp:       &idProvider{lock: &sync.Mutex{}},
		lock:      &sync.Mutex{},
	}
}

func (n *Notifier) Subscribe(ch chan []byte) int64 {
	id := n.idp.getNextID()
	n.lock.Lock()
	n.clientMap[id] = ch
	n.lock.Unlock()

	log.Println("register new client", id)

	return id
}

func (n *Notifier) Unsubscribe(id int64) {
	log.Println("deregister client", id)
	if ch, ok := n.clientMap[id]; ok {
		n.lock.Lock()
		delete(n.clientMap, id)
		n.lock.Unlock()
		close(ch)
	}
}

func (n *Notifier) Notify(data []byte) {
	for id, subscriber := range n.clientMap {
		go n.sendToSubscriber(id, subscriber, data)
	}
}

func (n *Notifier) NotifyOne(id int64, data []byte) {
	if subscriber, ok := n.clientMap[id]; ok {
		go n.sendToSubscriber(id, subscriber, data)
	} else {
		log.Printf("subscriber id %d was not found", id)
	}
}

func (n Notifier) sendToSubscriber(subID int64, subscriber chan<- []byte, data []byte) {
	subscriber <- data
	timeout := time.After(time.Millisecond * 100)
	select {
	case <-timeout:
		log.Printf("Failed timeout send a message timeout subscriber %d. Unbscribing...", subID)
		n.Unsubscribe(subID)
	default:
		return
	}
}
