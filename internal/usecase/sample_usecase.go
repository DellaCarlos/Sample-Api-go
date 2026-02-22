package usecase

import (
	"sample-api-go/internal/models"
	"sample-api-go/internal/repositories"
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
	sampleId, err := su.Repository.CreateSample(sample)
	if err != nil {
		return models.SampleModel{}, err
	}

	sample.ID = sampleId

	return sample, nil
}

func (su *SampleUseCase) SoftDeleteSampleByID(id_sample int) error {
	err := su.Repository.SoftDeleteSampleByID(id_sample)
	if err != nil {
		return err
	}

	return nil
}
