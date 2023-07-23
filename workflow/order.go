package workflow

import (
	"goenv/activity"
	"goenv/messages"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// start order workflow
func CreateOrderWorkflow(ctx workflow.Context, productsString string) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 20,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Minute,
			BackoffCoefficient: 2,
			MaximumAttempts:    7,
			MaximumInterval:    time.Hour,
			NonRetryableErrorTypes: []string{
				"activity_error",
			},
		},
	}

	// start workflow with activity options
	ctx = workflow.WithActivityOptions(ctx, options)

	// configure logger
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting CreateOrderWorkflow")

	// get running workflowId
	workflowId := workflow.GetInfo(ctx).WorkflowExecution.ID

	// result can be string, struct, other data types.
	var confirmInvoiceResult string
	var confirmShippingResult string

	// workflow selector
	invoiceSelector := workflow.NewSelector(ctx)
	shippingSelector := workflow.NewSelector(ctx)

	// workflow named signal channel
	invoiceSignalChan := workflow.GetSignalChannel(ctx, "confirmInvoice")
	shippingSignalChan := workflow.GetSignalChannel(ctx, "confirmShipping")

	// implement selector reciever via signal channel
	invoiceSelector.AddReceive(invoiceSignalChan, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &confirmInvoiceResult)
	})

	shippingSelector.AddReceive(shippingSignalChan, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &confirmInvoiceResult)
	})

	// start the activites

	// Create Order
	order := &messages.Order{}
	workflow.ExecuteActivity(ctx, activity.CreateOrder, productsString, workflowId).Get(ctx, order)

	// Create Transaction
	workflow.ExecuteActivity(ctx, activity.CreateTransaction, order.Id).Get(ctx, nil)

	// Create Shipping
	workflow.ExecuteActivity(ctx, activity.CreateShipping, order.Id).Get(ctx, nil)

	invoiceSelector.Select(ctx)

	// Confirm Transaction invoice
	workflow.ExecuteActivity(ctx, activity.ConfirmInvoice, order.Id, confirmInvoiceResult).Get(ctx, nil)

	shippingSelector.Select(ctx)

	// Confirm Shipping fulfillment
	workflow.ExecuteActivity(ctx, activity.ConfirmShipping, order.Id, confirmShippingResult).Get(ctx, nil)

	// Confirm Order
	workflow.ExecuteActivity(ctx, activity.ConfirmOrder, order.Id).Get(ctx, nil)

	return nil
}
