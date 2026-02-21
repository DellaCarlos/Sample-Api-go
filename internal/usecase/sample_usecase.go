package usecase

import (
	"sample-api-go/internal/models"
	"sample-api-go/internal/repositories"
	"time"
)

type SampleUseCase struct {
	Repository repositories.SampleRepository
}

func NewCreateSampleUseCase(repo repositories.SampleRepository) SampleUseCase {
	return SampleUseCase{
		Repository: repo,
	}
}

func (su *SampleUseCase) GetSamples() ([]models.SampleModel, error) {
	return su.Repository.GetSamples()
}

func (su *SampleUseCase) GetSampleByID(id_sample int) (*models.SampleModel, error) {
	sample, err := su.Repository.GetProductByID(id_sample)
	if err != nil {
		return nil, err
	}

	return sample, nil
}

func (su *SampleUseCase) CreateSample(sample models.SampleModel) (models.SampleModel, error) {

	now := time.Now() // tempo do momento para criar a amostra

	sampleId, err := su.Repository.CreateSample(sample)
	if err != nil {
		return models.SampleModel{}, err
	}

	sample.ID = sampleId
	sample.CreatedAt = now
	sample.UpdatedAt = now
	sample.DeletedAt = nil
	sample.IsActive = true

	return sample, nil
}
