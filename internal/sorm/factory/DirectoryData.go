package factory

import (
	sorm "github.com/alexdsayandexru/sorm/gen"
	"github.com/alexdsayandexru/sorm/internal/sorm/models"
)

func NewDirectoryData(s *sorm.DirectoryDataRequest) *models.DirectoryData {
	return &models.DirectoryData{
		EventType:     "directory",
		CorrelationId: s.CorrelationId,
		Datetime:      s.Datetime,
		OriServices:   ToServiceArray(s.Services),
	}
}

func ToServiceArray(sers []*sorm.Service) []models.Service {
	var result []models.Service
	for _, s := range sers {
		result = append(result, models.Service{
			ServiceId:   s.ServiceId,
			Description: s.Decription,
			BeginTime:   s.BeginTime,
			EndTime:     s.EndTime,
		})
	}
	return result
}
