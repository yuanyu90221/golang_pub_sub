package main

import (
	"context"
	"pub_sub/hub"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestSubscriber(t *testing.T) {
	ctx := context.Background()
	h := hub.NewHub()
	sub01 := hub.NewSubscriber("sub01")
	sub02 := hub.NewSubscriber("sub02")
	sub03 := hub.NewSubscriber("sub03")

	h.Subscribe(ctx, sub01)
	h.Subscribe(ctx, sub02)
	h.Subscribe(ctx, sub03)

	assert.Equal(t, 3, h.Subscribers())
	h.Unsubscribe(ctx, sub01)
	h.Unsubscribe(ctx, sub02)
	h.Unsubscribe(ctx, sub03)

	assert.Equal(t, 0, h.Subscribers())
}

func TestCancelSubscriber(t *testing.T) {
	ctx := context.Background()
	h := hub.NewHub()
	sub01 := hub.NewSubscriber("sub01")
	sub02 := hub.NewSubscriber("sub02")
	sub03 := hub.NewSubscriber("sub03")

	h.Subscribe(ctx, sub01)
	h.Subscribe(ctx, sub02)
	ctx03, cancel := context.WithCancel(ctx)
	h.Subscribe(ctx03, sub03)

	assert.Equal(t, 3, h.Subscribers())

	// cancel subscriber 03
	cancel()
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, 2, h.Subscribers())

	h.Unsubscribe(ctx, sub01)
	h.Unsubscribe(ctx, sub02)

	assert.Equal(t, 0, h.Subscribers())
}
