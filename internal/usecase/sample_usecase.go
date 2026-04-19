package usecase

import (
	apperrors "sample-api-go/internal/errors"
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
	sample, err := su.Repository.GetSampleByID(id_sample)
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
		return apperrors.Internal("invalid id")
	}
	return nil
}

func (su *SampleUseCase) HardDeleteSampleByID(id_sample int) error {
	err := su.Repository.HardDeleteSampleByID(id_sample)
	if err != nil {
		return apperrors.Internal("invalid id")
	}
	return nil
}

func (su *SampleUseCase) UpdateSample(id_sample int, input map[string]interface{}) (*models.SampleModel, error) {
	if isActive, ok := input["is_active_sample"]; ok {
		if isActive == false {
			input["deleted_at_sample"] = time.Now()
		} else {
			input["deleted_at_sample"] = nil
		}
	}

	sample, err := su.Repository.UpdateSample(id_sample, input)
	if err != nil {
		return nil, err
	}
	return sample, nil
}
