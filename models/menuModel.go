package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Menu struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty" `
	MenuId    string             `bson:"menuId" json:"menuId,omitempty" validate:"required"`
	Category  *string            `bson:"category" json:"category" validate:"required"`
	EndDate   *string            `bson:"endDate" json:"endDate" validate:"required"`
	StartDate *string            `bson:"startDate" json:"startDate" validate:"required"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt" validate:"required"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt" validate:"required"`
}
