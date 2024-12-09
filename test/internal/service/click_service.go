package service

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"test/internal/domain"
	"test/internal/dto"
)

type ClickRepo interface {
	GetStats(string, timestamppb.Timestamp, timestamppb.Timestamp) ([]domain.Click, error)
	Save(banner domain.Banner) error
	Update(string) error
}

type ClickService struct {
	clickRepo ClickRepo
}

func NewClickService(clickRepo ClickRepo) *ClickService {
	return &ClickService{clickRepo: clickRepo}
}

func (s *ClickService) GetStats(req dto.GetStatsRequest) (*dto.GetStatsResponse, error) {
	if err := req.Validaty(); err != nil {
		return nil, err
	}

	clicks, err := s.clickRepo.GetStats(req.BannerId, req.TsFrom, req.TsTo)
	if err != nil {
		return &dto.GetStatsResponse{}, err
	}

	resp := dto.GetStatsResponse{}
	resp = dto.NewGetStatsResponse(req.BannerId, req.TsFrom, req.TsTo, clicks)

	return &resp, nil
}

func (s *ClickService) Save(request dto.SaveRequest) (*dto.SaveResponse, error) {
	banner := domain.NewBanner(request.Title)
	if err := s.clickRepo.Save(*banner); err != nil {
		return &dto.SaveResponse{}, err
	}

	return &dto.SaveResponse{}, nil
}

func (s *ClickService) Update(request dto.UpdateRequest) (*dto.UpdateResponse, error) {
	if err := request.Validaty(); err != nil {
		return &dto.UpdateResponse{}, err
	}
	if err := s.clickRepo.Update(request.BannerId); err != nil {
		return &dto.UpdateResponse{}, err
	}

	return &dto.UpdateResponse{}, nil
}
