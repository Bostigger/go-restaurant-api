package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrderItem struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	OrderItemId *string            `bson:"orderItemId" json:"orderItemId" validate:"required"`
	OrderId     *string            `bson:"orderId" json:"orderId" validate:"required"`
	FoodId      *string            `bson:"foodId" json:"foodId" validate:"required"`
	Quantity    *int               `bson:"quantity" json:"quantity" validate:"required"`
	UnitPrice   *float64           `bson:"unitPrice" json:"unitPrice" validate:"required"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
}
