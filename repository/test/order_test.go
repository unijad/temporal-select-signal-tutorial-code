package test

import (
	"context"
	"goenv/messages"
	"goenv/repository"
	"testing"

	"github.com/google/uuid"
)

func TestOrder(t *testing.T) {
	r := &repository.Order{}
	workflowId := uuid.New().String()
	orderId := uuid.New().String()

	t.Run("CreateOrder", func(t *testing.T) {
		order, err := r.CreateOrder(context.Background(), &messages.Order{
			WorkflowId:     workflowId,
			Products:       []string{"1", "2"},
			InvoiceStatus:  "pending",
			ShippingStatus: "pending",
			OrderStatus:    "pending",
		})

		if err != nil {
			t.Fatal(err)
		}

		if order == nil {
			t.Fatal("order is nil")
		}

		if order.WorkflowId != workflowId {
			t.Fatal("order.WorkflowId is not " + workflowId)
		}

		orderId = order.Id
	})

	t.Run("GetOrder", func(t *testing.T) {
		order, err := r.GetOrder(context.Background(), orderId)
		if err != nil {
			t.Fatal(err)
		}

		if order == nil {
			t.Fatal("order is nil")
		}

		if order.WorkflowId != workflowId {
			t.Fatal("order.WorkflowId is not " + workflowId)
		}
	})

	t.Run("UpdateInvoiceStatus", func(t *testing.T) {
		err := r.UpdateInvoiceStatus(context.Background(), workflowId, "confirmed")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("UpdateShippingStatus", func(t *testing.T) {
		err := r.UpdateShippingStatus(context.Background(), workflowId, "confirmed")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("UpdateStatus", func(t *testing.T) {
		err := r.UpdateStatus(context.Background(), workflowId, "confirmed")
		if err != nil {
			t.Fatal(err)
		}
	})
}
