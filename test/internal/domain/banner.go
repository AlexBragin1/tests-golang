package domain

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Banner struct {
	ID     string  `json:"id" bson:"_id"`
	Title  string  `json:"title" bson:"title"`
	Clicks []Click `json:"clicks" bson:"clicks"`
}

type Click struct {
	CreatedAt timestamp.Timestamp `json:"created_at" bson:"created_at"`
}

func NewBanner(title string) *Banner {
	return &Banner{
		ID:    NewUUID(),
		Title: title,
	}
}
