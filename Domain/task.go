package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`
	Title string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	DueDate time.Time `bson:"due_date" json:"due_date"`
	Status string `bson:"status" json:"status"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}