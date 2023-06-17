package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Invoice struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id"`
	InvoiceId      *string            `bson:"invoiceId" json:"invoiceId" validate:"required"`
	OrderId        *string            `bson:"orderId" json:"orderId" validate:"required"`
	PaymentMethod  *string            `bson:"paymentMethod" json:"paymentMethod" validate:"required"`
	PaymentStatus  *string            `bson:"paymentStatus" json:"paymentStatus" validate:"required"`
	PaymentDueDate time.Time          `bson:"paymentDueDate" json:"paymentDueDate" validate:"required"`
	CreatedAt      time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt" json:"updatedAt"`
}
