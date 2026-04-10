package models

import "github.com/google/uuid"

type SectorModel struct {
	ID     uuid.UUID `json:"id_sector"`
	Sector string    `json:"name"`
}
