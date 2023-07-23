package workflow

import (
	"goenv/activity"
	"goenv/messages"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// define the workflow function
func SetCartWorkflow(ctx workflow.Context, cart *messages.Cart) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// start the activities
	result := &messages.Cart{}
	err := workflow.ExecuteActivity(ctx, activity.SetCart, cart).Get(ctx, result)
	if err != nil {
		return temporal.NewApplicationError(err.Error(), "error")
	}

	return nil
}

// define the workflow function
func GetCartWorkflow(ctx workflow.Context) (*[]messages.Product, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// start the activities
	result := &[]messages.Product{}
	err := workflow.ExecuteActivity(ctx, activity.GetCart).Get(ctx, result)
	if err != nil {
		return nil, temporal.NewApplicationError(err.Error(), "error")
	}

	return result, nil
}
