package messages

import "github.com/lib/pq"

type Product struct {
	Name  string
	Price float64
}

type Cart struct {
	Products pq.StringArray `gorm:"type:text[]"` // Product.Id
}
