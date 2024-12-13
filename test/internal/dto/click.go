package dto

import (
	"errors"
	"fmt"
	"time"

	"test/internal/domain"
)

type GetStatsRequest struct {
	BannerId string    `json:"banner_id" `
	TsFrom   time.Time `json:"ts_from"`
	TsTo     time.Time `json:"ts_to"`
}
type GetStatsResponse struct {
	Stats []domain.Stat `json:"stats"`
}

type UpdateRequest struct {
	ID string `json:"id" bson:"_id"`
}

type UpdateResponse struct{}

type SaveRequest struct {
	Title string `json:"title" bson:"title"`
}
type SaveResponse struct{}

func (d SaveRequest) Validaty() error {
	if d.Title == "" {
		return errors.New("title is required")
	}

	return nil
}
func (d UpdateRequest) Validaty() error {
	if d.ID == "" {
		return errors.New("banner_id is required")
	}

	return nil
}

func (d GetStatsRequest) Validaty() error {
	if d.TsFrom.IsZero() {
		return fmt.Errorf("must provide timestamp")
	}
	if d.TsTo.IsZero() {
		return fmt.Errorf("must provide timestamp")
	}

	return nil
}
