package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	NoteMessage *string            `bson:"noteMessage" json:"noteMessage" validate:"required"`
	NoteId      *string            `bson:"noteId" json:"noteId" validate:"required"`
	CreatedAt   string             `bson:"createdAt" json:"createdAt" validate:"required"`
	UpdateAt    string             `bson:"updateAt" json:"updateAt"`
}
