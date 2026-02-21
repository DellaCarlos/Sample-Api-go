package repositories

import (
	"database/sql"
	"fmt"
	"sample-api-go/internal/models"

	"github.com/lib/pq"
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

func (sr *SampleRepository) GetSamples() ([]models.SampleModel, error) {
	query := `
		SELECT 
			id_sample,
			name_sample,
			sector_sample,
			analysis_sample,
			created_by_user_id_sample,
			created_at_sample,
			updated_at_sample,
			deleted_at_sample,
			is_active_sample
		FROM samples
	`

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
			pq.Array(&sampleObj.Analysis),
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

func (sr *SampleRepository) GetProductByID(id_sample int) (*models.SampleModel, error) {
	query, err := sr.connection.Prepare("SELECT * FROM samples WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var sample models.SampleModel
	err = query.QueryRow(id_sample).Scan(
		&sample.ID,
		&sample.Name,
		&sample.Sector,
		&sample.Analysis,
		&sample.CreatedByUserID,
		&sample.CreatedAt,
		&sample.UpdatedAt,
		&sample.DeletedAt,
		&sample.IsActive,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &sample, nil
}

func (sr *SampleRepository) CreateSample(sample models.SampleModel) (int, error) {
	var id int

	query, err := sr.connection.Prepare(
		"INSERT INTO samples (" +
			"name_sample, sector_sample, analysis_sample, created_by_user_id_sample, created_at_sample, updated_at_sample, deleted_at_sample, is_active_sample" +
			") VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id_sample",
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
		sample.CreatedAt,
		sample.UpdatedAt,
		sample.DeletedAt,
		sample.IsActive,
	).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}
