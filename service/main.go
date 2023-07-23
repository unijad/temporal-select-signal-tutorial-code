package service

import (
	"goenv/activity"
	"goenv/workflow"
	"log"
	"net/http"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func Server() {
	// set up the worker
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "cart", worker.Options{})
	w.RegisterWorkflow(workflow.CreateOrderWorkflow)
	w.RegisterActivity(activity.CreateOrder)
	w.RegisterActivity(activity.CreateTransaction)
	w.RegisterActivity(activity.CreateShipping)
	w.RegisterActivity(activity.ConfirmShipping)
	w.RegisterActivity(activity.ConfirmInvoice)
	w.RegisterActivity(activity.CreateInvoice)
	w.RegisterActivity(activity.ConfirmOrder)

	mux := http.NewServeMux()
	mux.HandleFunc("/cart", CartGetHandler)       // curl -X GET http://localhost:5000/cart
	mux.HandleFunc("/cart/set", CartSetHandler)   // curl -X POST http://localhost:5000/cart/set\?products\=1,2,3
	mux.HandleFunc("/order/create", OrderHandler) // curl -X POST http://localhost:5000/order/create?products=1,2,3
	// curl -X POST http://localhost:5000/order/signal?orderId={orderID}&signalName=confirmInvoice&status=confirmed
	// curl -X POST http://localhost:5000/order/signal?orderId={orderID}&signalName=confirmShipping&status=confirmed
	mux.HandleFunc("/order/signal", SignalOrderHandler)

	server := &http.Server{Addr: ":5000", Handler: mux}

	// start the worker and the web server
	go func() {
		err = w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalln("unable to start Worker", err)
		}
	}()

	log.Fatal(server.ListenAndServe())
}
