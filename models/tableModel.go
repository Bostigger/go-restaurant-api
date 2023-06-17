package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Table struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id"`
	TableName      *string            `bson:"tableName" json:"tableName" validate:"required"`
	TableType      *string            `bson:"tableType" json:"tableType"`
	TableNumber    *int               `bson:"tableNumber" json:"tableNumber"`
	NumberOfGuests *int               `bson:"numberOfGuests" json:"numberOfGuests"`
	TableId        *string            `bson:"tableId" json:"tableId"`
	CreatedAt      time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt" json:"updatedAt"`
}
