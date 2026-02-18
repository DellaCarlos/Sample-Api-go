package repositories

import "sample-api-go/internal/domain/models"

type ISampleRepository interface {
	Save(sampple *models.SampleModel) error
}
