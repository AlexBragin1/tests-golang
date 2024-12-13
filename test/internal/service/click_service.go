package service

import (
	"time"

	"test/internal/domain"
	"test/internal/dto"
)

type ClickRepo interface {
	GetStats(string, time.Time, time.Time) ([]domain.Stat, error)
	Save(banner domain.Banner) error
	Update(string, domain.Click) error
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

	stats, err := s.clickRepo.GetStats(req.BannerId, req.TsFrom, req.TsTo)
	if err != nil {
		return &dto.GetStatsResponse{Stats: stats}, err
	}

	return &dto.GetStatsResponse{Stats: stats}, err
}

func (s *ClickService) Save(request dto.SaveRequest) (*dto.SaveResponse, error) {
	banner := domain.NewBanner(request.Title)
	if err := s.clickRepo.Save(*banner); err != nil {
		return &dto.SaveResponse{}, err
	}

	return &dto.SaveResponse{}, nil
}

func (s *ClickService) Update(request dto.UpdateRequest) (*dto.UpdateResponse, error) {
	click := domain.Click{CreatedAt: time.Now()}

	if err := s.clickRepo.Update(request.ID, click); err != nil {
		return &dto.UpdateResponse{}, err
	}

	return &dto.UpdateResponse{}, nil
}
