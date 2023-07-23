package activity

import (
	"context"
	"goenv/messages"
	"goenv/repository"
)

func SetCart(ctx context.Context, cart *messages.Cart) (err error) {
	c := &repository.Cart{}

	err = c.SetCart(ctx, cart)
	if err != nil {
		return err
	}

	return err
}

func GetCart(ctx context.Context) (data *[]messages.Product, err error) {
	c := &repository.Cart{}

	data, err = c.GetCart(ctx)
	if err != nil {
		return nil, err
	}

	return data, err
}
