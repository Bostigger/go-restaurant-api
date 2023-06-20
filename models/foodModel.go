package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Food struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      *string            `bson:"name" json:"name" validate:"required"`
	FoodImage *string            `bson:"foodImage" validate:"required"`
	Price     *float64           `bson:"price" json:"price" validate:"required"`
	FoodId    string             `bson:"foodId" json:"foodId"`
	MenuId    *string            `bson:"menuId" json:"menuId"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" bson:"updatedAt"`
}
