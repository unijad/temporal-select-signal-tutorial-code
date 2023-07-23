package activity

import (
	"context"
	"errors"
	"goenv/messages"
	"goenv/repository"

	"github.com/lib/pq"
	"go.temporal.io/sdk/temporal"
)

func CreateOrder(ctx context.Context, productsString string, workflowId string) (*messages.Order, error) {
	c := &repository.Order{}
	order, err := c.CreateOrder(ctx, &messages.Order{
		WorkflowId:  workflowId,
		Products:    pq.StringArray{productsString},
		OrderStatus: "pending",
	})
	if err != nil {
		return nil, temporal.NewNonRetryableApplicationError(err.Error(), "activity_error", err)
	}
	return order, nil
}

func CreateTransaction(ctx context.Context, w45f3t45rforkflowId string) (err error) {
	c := &repository.Order{}
	// create transaction record
	print("CreateTransaction", c)
	return err
}

func CreateShipping(ctx context.Context, workflowId string) (err error) {
	c := &repository.Order{}
	// create shipping record
	print("CreateShipping", c)
	return err
}

func ConfirmShipping(ctx context.Context, workflowId string, status string) (err error) {
	c := &repository.Order{}

	err = c.UpdateShippingStatus(ctx, workflowId, status)
	if err != nil {
		return temporal.NewNonRetryableApplicationError(err.Error(), "activity_error", err)
	}

	return err
}

func ConfirmOrder(ctx context.Context, workflowId string, status string) (err error) {
	c := &repository.Order{}

	err = c.UpdateStatus(ctx, workflowId, status)
	if err != nil {
		return temporal.NewNonRetryableApplicationError(err.Error(), "activity_error", err)
	}

	return err
}

func CreateInvoice(ctx context.Context, workflowId string) (err error) {
	c := &repository.Order{}
	// keep retying until shipping is confirmed
	print("CreateInvoice", c)

	err = temporal.NewApplicationError(errors.New("status pending").Error(), "waiting_status_change")

	return err
}

func ConfirmInvoice(ctx context.Context, workflowId string, status string) (err error) {
	c := &repository.Order{}

	err = c.UpdateInvoiceStatus(ctx, workflowId, status)
	if err != nil {
		return temporal.NewNonRetryableApplicationError(err.Error(), "activity_error", err)
	}

	return err
}
