package test

import (
	"context"
	"goenv/activity"
	"goenv/messages"
	"testing"
)

func TestCart(t *testing.T) {
	ctx := context.Background()

	t.Run("SetCart", func(t *testing.T) {
		err := activity.SetCart(ctx, &messages.Cart{
			Products: []string{"1", "2"},
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetCart", func(t *testing.T) {
		cart, err := activity.GetCart(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if cart == nil {
			t.Fatal("record is nil")
		}
		if len(*cart) == 0 {
			t.Fatal("invalid length")
		}
	})
}
