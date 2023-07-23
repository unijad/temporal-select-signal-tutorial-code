package service

import (
	"encoding/json"
	"goenv/repository"
	"goenv/workflow"
	"log"
	"net/http"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	productsString := r.URL.Query().Get("products")
	// create a new temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	_, err = c.ExecuteWorkflow(r.Context(), client.StartWorkflowOptions{
		ID:        "CreateOrderWorkflow_" + uuid.New().String(),
		TaskQueue: "cart",
	}, workflow.CreateOrderWorkflow, productsString)
	if err != nil {
		http.Error(w, "unable to start workflow", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal("{status: 'ok'}")
	if err != nil {
		http.Error(w, "unable to marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}

func SignalOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderId := r.URL.Query().Get("orderId")
	signalName := r.URL.Query().Get("signal")
	orderStatus := r.URL.Query().Get("status")
	// create a new temporal client
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// get order from data layer
	repo := &repository.Order{}
	order, err := repo.GetOrder(r.Context(), orderId)
	if err != nil {
		http.Error(w, "unable to get order", http.StatusInternalServerError)
		return
	}

	workflowId := order.WorkflowId
	runId := "" // we did not store runId we can safely leave it empty
	err = c.SignalWorkflow(r.Context(), workflowId, runId, signalName, orderStatus)
	if err != nil {
		http.Error(w, "unable to signal workflow", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal("{status: 'ok'}")
	if err != nil {
		http.Error(w, "unable to marshal response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
