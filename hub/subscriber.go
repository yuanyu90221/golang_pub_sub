package hub

import (
	"context"
	"log"
	"sync"
)

type Message struct {
	Data []byte
}

type subscriber struct {
	sync.Mutex

	name    string
	handler chan *Message
	quit    chan struct{}
}

func (s *subscriber) run(ctx context.Context) {
	for {
		select {
		case msg := <-s.handler:
			log.Println(s.name, string(msg.Data))
		case <-s.quit:
			return
		case <-ctx.Done():
			return
		}
	}
}

func NewSubscriber(name string) *subscriber {
	// log.Println(name)
	return &subscriber{
		name:    name,
		handler: make(chan *Message, 100),
		quit:    make(chan struct{}),
	}
}

func (s *subscriber) publish(ctx context.Context, msg *Message) {
	select {
	case <-ctx.Done():
		return
	case s.handler <- msg:
	default:
	}
}
