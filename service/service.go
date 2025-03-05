package service

import (
	"fmt"
	"inMemoryDb/repository"
	"time"
)

type Service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) CreateData(name string) (*repository.Data, error) {
	newData := &repository.Data{Name: name}

	err := s.repository.Add(*newData, 1*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("veri oluşturulamadı: %v", err)
	}

	return newData, nil
}

func (s *Service) GetAllData() ([]repository.Data, error) {
	allData, err := s.repository.GetAll()
	if err != nil {
		return nil, fmt.Errorf("veri alınamadı: %v", err)
	}

	return allData, nil
}
