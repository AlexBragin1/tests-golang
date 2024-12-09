package dto

import (
	"errors"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"test/internal/domain"
)

type GetStatsRequest struct {
	BannerId string                `json:"banner_id" `
	TsFrom   timestamppb.Timestamp `json:"ts_from"`
	TsTo     timestamppb.Timestamp `json:"ts_to"`
}
type GetStatsResponse struct {
	BannerId string    `json:"banner_id"`
	TsFrom   time.Time `json:"ts_from"`
	TsTo     time.Time `json:"ts_to"`
	Count    int64     `json:"count"`
}

type UpdateRequest struct {
	BannerId string `json:"banner_id" `
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
	if d.BannerId == "" {
		return errors.New("banner_id is required")
	}

	return nil
}
func NewGetStatsResponse(bannerID string, tsFrom timestamppb.Timestamp, tsTo timestamppb.Timestamp, clicks []domain.Click) GetStatsResponse {
	return GetStatsResponse{
		BannerId: bannerID,
		TsFrom:   tsFrom.AsTime(),
		TsTo:     tsTo.AsTime(),
		Count:    int64(len(clicks)),
	}
}

func (d GetStatsRequest) Validaty() error {
	if d.TsFrom.IsValid() {
		return fmt.Errorf("must provide timestamp")
	}
	if d.TsTo.IsValid() {
		return fmt.Errorf("must provide timestamp")
	}
	tsFrom := time.Date(2023, time.October, 10, 10, 0, 0, 0, time.UTC)
	tsTo := time.Date(2023, time.October, 10, 10, 1, 0, 0, time.UTC)

	if !tsFrom.IsZero() {
		duration := tsTo.Sub(tsFrom)
		if duration != time.Minute {
			errors.New("difference between tsFrom and tsTo not equal one minute")
		}
	}
	return nil
}
