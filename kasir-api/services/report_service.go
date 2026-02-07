package services

import (
	"kasir-api/repositories"
	"kasir-api/models"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{ repo:repo }
}


func (s *ReportService) GetSalesSummary(startDate string, endDate string)(*models.SalesSummary, error){
	return s.repo.GetSummary(startDate, endDate)
}