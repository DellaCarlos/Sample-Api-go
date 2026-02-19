package repositories

import (
	"database/sql"
	"fmt"
	"sample-api-go/internal/domain/models"
	"time"
)

type SampleRepository struct {
	connection *sql.DB
}

// NewSampleRepository creates a new instance of SampleRepository with the provided database connection
func NewSampleRepository(connection *sql.DB) SampleRepository {
	return SampleRepository{
		connection: connection,
	}
}

func (sr *SampleRepository) GetSample() ([]models.SampleModel, error) {

	query := "SELECT * FROM samples"
	rows, err := sr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.SampleModel{}, err
	}

	var sampleList []models.SampleModel
	var sampleObj models.SampleModel

	for rows.Next() {
		err = rows.Scan(
			&sampleObj.ID,
			&sampleObj.Name,
			&sampleObj.Sector,
			&sampleObj.Analysis,
			&sampleObj.CreatedByUserID,
			&sampleObj.CreatedAt,
			&sampleObj.UpdatedAt,
			&sampleObj.DeletedAt,
			&sampleObj.IsActive,
		)

		if err != nil {
			fmt.Println(err)
			return []models.SampleModel{}, err
		}

		sampleList = append(sampleList, sampleObj)
	}

	rows.Close()
	return sampleList, nil
}

func (sr *SampleRepository) CreateSample(sample models.SampleModel) (int, error) {
	var id int
	now := time.Now()

	query, err := sr.connection.Prepare(
		"INSERT INTO samples (name, sector, analysis, created_by_user_id, created_at, updated_at, deleted_at, is_active)" +
			" VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
	)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(
		sample.Name,
		sample.Sector,
		sample.Analysis,
		sample.CreatedByUserID,
		now,  // created_at
		now,  // updated_at
		nil,  // deleted_at
		true, // is_active
	).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}
