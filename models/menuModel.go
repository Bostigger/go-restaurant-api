package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Menu struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	MenuId    *string            `bson:"menuId" json:"menuId"`
	FoodId    *string            `bson:"foodId" json:"foodId"`
	Category  *string            `bson:"category" json:"category"`
	EndDate   *string            `bson:"endDate" json:"endDate"`
	StartDate *string            `bson:"startDate" json:"startDate"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
