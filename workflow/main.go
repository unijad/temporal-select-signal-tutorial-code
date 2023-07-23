package workflow

import (
	"goenv/activity"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func WorkflowDifnition() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "cart", worker.Options{})

	// register workflows
	w.RegisterWorkflow(GetCartWorkflow)
	w.RegisterWorkflow(SetCartWorkflow)
	w.RegisterWorkflow(CreateOrderWorkflow)

	// register activities
	w.RegisterActivity(activity.GetCart)
	w.RegisterActivity(activity.SetCart)
	w.RegisterActivity(activity.CreateOrder)
	w.RegisterActivity(activity.ConfirmShipping)
	w.RegisterActivity(activity.ConfirmInvoice)
	w.RegisterActivity(activity.ConfirmOrder)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
