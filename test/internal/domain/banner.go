package domain

import (
	"time"
)

type Banner struct {
	ID     string  `json:"id" bson:"_id"`
	Title  string  `json:"title" bson:"title"`
	Clicks []Click `json:"clicks" bson:"clicks"`
}

type Click struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func NewBanner(title string) *Banner {
	return &Banner{
		ID:     NewID(),
		Title:  title,
		Clicks: []Click{},
	}
}
