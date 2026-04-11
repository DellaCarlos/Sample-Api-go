package repositories

import (
	"database/sql"
	"sample-api-go/internal/models"

	"github.com/google/uuid"
)

type SectorRepository struct {
	connection *sql.DB
}

// NewSampleRepository creates a new instance of SampleRepository
// with the provided database connection
func NewSectorRepository(connection *sql.DB) SectorRepository {
	return SectorRepository{
		connection: connection,
	}
}

func (scr *SectorRepository) GetSectors() ([]models.SectorModel, error) {
	query := "SELECT id_sector, name FROM sectors"

	rows, err := scr.connection.Query(query)
	if err != nil {
		return []models.SectorModel{}, err
	}

	var sectorList []models.SectorModel
	var sectorObj models.SectorModel

	for rows.Next() {
		err = rows.Scan(
			&sectorObj.ID,
			&sectorObj.Sector,
		)

		if err != nil {
			return []models.SectorModel{}, err
		}

		sectorList = append(sectorList, sectorObj)
	}

	rows.Close()
	return sectorList, nil
}

func (scr *SectorRepository) CreateSector(sector models.SectorModel) (uuid.UUID, error) {
	id := uuid.New()
	query, err := scr.connection.Prepare(
		"INSERT INTO sectors (id_sector, name) VALUES ($1, $2)",
	)

	if err != nil {
		return uuid.Nil, err
	}
	defer query.Close()

	_, err = query.Exec(id, sector.Sector)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
