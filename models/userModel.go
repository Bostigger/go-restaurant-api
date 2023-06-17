package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id"`
	UserId       string             `bson:"userId" json:"userId"`
	Username     string             `bson:"username" json:"username" validate:"required"`
	Password     string             `bson:"password" json:"password" validate:"required"`
	UserType     string             `bson:"userType" json:"userType" validate:"required"`
	PhoneNumber  string             `bson:"phoneNumber" json:"phoneNumber" validate:"required"`
	Email        string             `bson:"email" json:"email" validate:"email,required"`
	Token        string             `bson:"token'" json:"token"`
	RefreshToken string             `bson:"refreshToken" json:"refreshToken"`
	CreatedAt    time.Time          `json:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}
