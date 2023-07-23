package messages

import "github.com/lib/pq"

type Order struct {
	Id             string `gorm:"primary_key"`
	WorkflowId     string
	Products       pq.StringArray `gorm:"type:text[]"` // []Product.Id
	InvoiceStatus  string         `gorm:"type:text[]"` // "pending", "confirmed"
	ShippingStatus string         `gorm:"type:text[]"` // "pending", "confirmed"
	OrderStatus    string         `gorm:"type:text[]"` // "pending", "confirmed"
}

type OrderWorkflowSignal struct {
	WorkflowId string
	Status     string
}
