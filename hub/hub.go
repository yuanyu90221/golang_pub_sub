package hub

import (
	"context"
	"sync"
)

type Hub struct {
	sync.Mutex
	subs map[*subscriber]struct{}
}

func NewHub() *Hub {
	return &Hub{
		subs: map[*subscriber]struct{}{},
	}
}

// 訂閱 新增 subscriber
func (h *Hub) Subscribe(ctx context.Context, s *subscriber) error {
	h.Lock()
	h.subs[s] = struct{}{}
	h.Unlock()

	// handler context cancel
	go func() {
		select {
		case <-s.quit:
		case <-ctx.Done(): // cancel
			h.Lock()
			delete(h.subs, s)
			h.Unlock()
		}
	}()
	go s.run(ctx)
	return nil
}

func (h *Hub) Unsubscribe(ctx context.Context, s *subscriber) error {
	h.Lock()
	delete(h.subs, s)
	h.Unlock()
	close(s.quit)
	return nil
}

func (h *Hub) Publish(ctx context.Context, msg *Message) error {
	h.Lock()
	for s := range h.subs {
		s.publish(ctx, msg)
	}
	h.Unlock()
	return nil
}

func (h *Hub) Subscribers() int {
	return len(h.subs)
}
