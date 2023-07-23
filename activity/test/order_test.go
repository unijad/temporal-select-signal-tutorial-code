package test

import (
	"context"
	"goenv/activity"
	"testing"

	"github.com/google/uuid"
)

func TestOrder(t *testing.T) {
	ctx := context.Background()
	workflowId := uuid.New().String()

	// CreateOrder
	t.Run("CreateOrder", func(t *testing.T) {
		order, err := activity.CreateOrder(ctx, "1,2,3", workflowId)
		if err != nil {
			t.Fatal(err)
		}

		if order == nil {
			t.Fatal("order is nil")
		}
	})

	// ConfirmInvoice
	t.Run("ConfirmInvoice", func(t *testing.T) {
		err := activity.ConfirmInvoice(ctx, workflowId, "confirmed")
		if err != nil {
			t.Fatal(err)
		}
	})

	// ConfirmShipping
	t.Run("ConfirmShipping", func(t *testing.T) {
		err := activity.ConfirmShipping(ctx, workflowId, "confirmed")
		if err != nil {
			t.Fatal(err)
		}
	})
}
