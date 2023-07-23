package repository

import (
	"context"
	"goenv/messages"

	"github.com/google/uuid"
)

type Order struct{}

func (o *Order) GetOrder(ctx context.Context, orderId string) (order *messages.Order, err error) {
	db := &Repository{}
	err = db.Connect()
	if err != nil {
		return order, err
	}

	// get order by id
	order = &messages.Order{}
	records := db.gorm.Where("id = ?", orderId).First(&order)
	if records.Error != nil {
		return order, records.Error
	}

	return order, err
}

func (o *Order) CreateOrder(ctx context.Context, order *messages.Order) (*messages.Order, error) {
	db := &Repository{}
	err := db.Connect()
	if err != nil {
		return nil, err
	}

	order.Id = uuid.New().String()

	records := db.gorm.Create(&order)
	if records.Error != nil {
		return nil, records.Error
	}

	return order, nil
}

func (o *Order) UpdateInvoiceStatus(ctx context.Context, workflowId string, status string) (err error) {
	db := &Repository{}
	err = db.Connect()
	if err != nil {
		return err
	}

	// update order by id, invoiceStatus
	records := db.gorm.Model(&messages.Order{}).Where("workflow_id = ?", workflowId).Update("invoice_status", status)
	if records.Error != nil {
		return records.Error
	}

	return err
}

func (o *Order) UpdateShippingStatus(ctx context.Context, workflowId string, status string) (err error) {
	db := &Repository{}
	err = db.Connect()
	if err != nil {
		return err
	}

	// update order by id, invoiceStatus
	records := db.gorm.Model(&messages.Order{}).Where("workflow_id = ?", workflowId).Update("shipping_status", status)
	if records.Error != nil {
		return records.Error
	}

	return err
}

func (o *Order) UpdateStatus(ctx context.Context, workflowId string, status string) (err error) {
	db := &Repository{}
	err = db.Connect()
	if err != nil {
		return err
	}

	// update order by id, invoiceStatus
	records := db.gorm.Model(&messages.Order{}).Where("workflow_id = ?", workflowId).Update("order_status", status)
	if records.Error != nil {
		return records.Error
	}

	return err
}
