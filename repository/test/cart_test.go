package test

import (
	"context"
	"goenv/messages"
	"goenv/repository"
	"testing"
)

func TestCart(t *testing.T) {
	r := &repository.Cart{}

	t.Run("SetCart", func(t *testing.T) {
		err := r.SetCart(context.Background(), &messages.Cart{
			Products: []string{"1", "2"},
		})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("GetCart", func(t *testing.T) {
		cart, err := r.GetCart(context.Background())
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

	t.Run("GetProducts", func(t *testing.T) {
		products, err := r.GetProducts(context.Background(), &messages.Cart{
			Products: []string{"1", "2"},
		})
		if err != nil {
			t.Fatal(err)
		}
		if products == nil {
			t.Fatal("record is nil")
		}

		if len(*products) == 0 {
			t.Fatal("invalid length")
		}
	})
}
