package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID
	Invoice_id       string
	Order_id         string
	Payment_method   *string
	Payment_status   *string
	Payment_due_date time.Time
	Created_at       time.Time
	updated_at       time.Time
}
