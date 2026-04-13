package usecase

import (
	"sample-api-go/internal/models"
	"sample-api-go/internal/repositories"
	"strings"

	apperrors "sample-api-go/internal/errors"
)

type SectorUseCase struct {
	Repository repositories.SectorRepository
}

func NewSectorUseCase(repo repositories.SectorRepository) SectorUseCase {
	return SectorUseCase{
		Repository: repo,
	}
}

func (scu *SectorUseCase) GetSectors() ([]models.SectorModel, error) {
	return scu.Repository.GetSectors()
}

func (scu *SectorUseCase) CreateSector(sector models.SectorModel) (models.SectorModel, error) {
	if strings.TrimSpace(sector.Sector) == "" {
		return sector, apperrors.BadRequest("invalid sector.")
	}

	id, err := scu.Repository.CreateSector(sector)
	if err != nil {
		return models.SectorModel{}, err
	}

	sector.ID = id

	return sector, nil
}
