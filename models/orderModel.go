package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id" validate:"required"`
	OderId    *string            `bson:"oderId" json:"orderId" validate:"required"`
	OrderDate *string            `bson:"orderDate" json:"orderDate" validate:"required"`
	TableId   *string            `bson:"tableId" json:"tableId" validate:"required"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt" validate:"required"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
