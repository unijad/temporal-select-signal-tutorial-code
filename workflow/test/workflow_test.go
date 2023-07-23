package test

import (
	"goenv/activity"
	"goenv/messages"
	"goenv/workflow"

	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestWorkflows(t *testing.T) {
	// Set up the test suite and testing execution environment
	testSuite := &testsuite.WorkflowTestSuite{}

	t.Run("SetCart", func(t *testing.T) {
		env := testSuite.NewTestWorkflowEnvironment()
		// Mock activity implementation
		env.OnActivity(activity.SetCart, mock.Anything, mock.Anything).Return(nil)
		env.ExecuteWorkflow(workflow.SetCartWorkflow)

		// Verify that the SetCart activity was executed
		require.NoError(t, env.GetWorkflowError())
		require.True(t, env.IsWorkflowCompleted())
	})

	t.Run("GetCart", func(t *testing.T) {
		env := testSuite.NewTestWorkflowEnvironment()
		// Mock activity implementation
		products := &[]messages.Product{
			{
				Name:  "Product 1",
				Price: 1.1,
			},
			{
				Name:  "Product 2",
				Price: 1.1,
			},
			{
				Name:  "Product 3",
				Price: 1.1,
			},
		}
		env.OnActivity(activity.GetCart, mock.Anything).Return(products, nil)
		env.ExecuteWorkflow(workflow.GetCartWorkflow)

		require.True(t, env.IsWorkflowCompleted())
		require.NoError(t, env.GetWorkflowError())

		var data *[]messages.Product
		require.NoError(t, env.GetWorkflowResult(&data))
		require.Equal(t, products, data)
	})

	t.Run("CreateOrderWorkflow", func(t *testing.T) {
		env := testSuite.NewTestWorkflowEnvironment()

		env.OnActivity(activity.CreateOrder, mock.Anything, mock.Anything).Return(nil)
		env.OnActivity(activity.CreateTransaction, mock.Anything, mock.Anything).Return(nil)
		env.OnActivity(activity.CreateShipping, mock.Anything, mock.Anything).Return(nil)
		env.OnActivity(activity.CreateInvoice, mock.Anything, mock.Anything).Return(nil)

		// execute in go-routine to avoid blocking
		env.ExecuteWorkflow(workflow.CreateOrderWorkflow, &messages.Cart{
			Products: []string{"1", "2"},
		})

		// wait for the first activity to complete
		env.AssertExpectations(t)
		require.True(t, env.IsWorkflowCompleted())
	})
}
