package repository

import (
	"context"
	"goenv/messages"

	_ "github.com/lib/pq"
)

type Cart struct{}

func (w *Cart) GetCart(ctx context.Context) (products *[]messages.Product, err error) {
	db := &Repository{}
	err = db.Connect()
	if err != nil {
		return nil, err
	}

	cart := &messages.Cart{}
	// get all records and assign to data
	records := db.gorm.Last(&cart)
	if records.Error != nil {
		return nil, records.Error
	}

	products, err = w.GetProducts(ctx, cart)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (w *Cart) SetCart(ctx context.Context, cart *messages.Cart) error {
	db := &Repository{}
	err := db.Connect()
	if err != nil {
		return err
	}

	records := db.gorm.Create(&cart)
	if records.Error != nil {
		return records.Error
	}

	return nil
}

func (w *Cart) GetProducts(ctx context.Context, cart *messages.Cart) (*[]messages.Product, error) {
	db := &Repository{}
	err := db.Connect()
	if err != nil {
		return nil, err
	}

	data := []messages.Product{}
	// get all records and assign to data
	records := db.gorm.Find(&data)
	if records.Error != nil {
		return nil, records.Error
	}
	return &data, nil
}
