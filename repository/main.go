package repository

import (
	"goenv/messages"

	_ "github.com/lib/pq"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DB_USER     = "1"
	DB_PASSWORD = ""
	DB_NAME     = "temporal_example"
)

type Repository struct {
	gorm *gorm.DB
}

func (d *Repository) Connect() error {
	db, err := gorm.Open(sqlite.Open("store.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&messages.Cart{}, &messages.Product{}, &messages.Order{})
	if err != nil {
		panic("failed to migrate database")
	}

	var count int64
	db.Model(&messages.Product{}).Count(&count)
	if count == 0 {
		err := db.Create([]messages.Product{
			{Name: "Product 1", Price: 10.0},
			{Name: "Product 2", Price: 20.0},
			{Name: "Product 3", Price: 30.0},
		})
		if err.Error != nil {
			panic("failed to create products")
		}
	}

	d.gorm = db

	return nil
}
