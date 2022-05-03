package main

import (
	"context"
	"time"

	"pub_sub/hub"
)

func main() {
	ctx := context.Background()
	h := hub.NewHub()
	sub01 := hub.NewSubscriber("sub01")
	sub02 := hub.NewSubscriber("sub02")
	sub03 := hub.NewSubscriber("sub03")

	h.Subscribe(ctx, sub01)
	h.Subscribe(ctx, sub02)
	h.Subscribe(ctx, sub03)
	_ = h.Publish(ctx, &hub.Message{Data: []byte("test01")})
	_ = h.Publish(ctx, &hub.Message{Data: []byte("test02")})
	_ = h.Publish(ctx, &hub.Message{Data: []byte("test03")})
	time.Sleep(1 * time.Second)

	h.Unsubscribe(ctx, sub03)
	_ = h.Publish(ctx, &hub.Message{Data: []byte("test04")})
	_ = h.Publish(ctx, &hub.Message{Data: []byte("test05")})
	time.Sleep(1 * time.Second)
}
