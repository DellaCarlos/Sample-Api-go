package models

import "github.com/google/uuid"

type SectorModel struct {
	ID       uuid.UUID `json:"id_sector"`
	SectorID int       `json: sector_id`
	Sector   string    `json:"sector_name"`
}
