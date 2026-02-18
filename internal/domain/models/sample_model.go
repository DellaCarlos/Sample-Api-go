package models

import "time"

// SampleModel is a struct that represents a sample data model

type SampleModel struct {
	ID              int        `json:"id_sample"`
	Name            string     `json:"name_sample"`
	Sector          string     `json:"sector_sample"`
	Analysis        []string   `json:"analysis_sample"`
	CreatedByUserID int        `json:"created_by_user_id_sample"`
	CreatedAt       time.Time  `json:"created_at_sample"`
	UpdatedAt       time.Time  `json:"updated_at_sample"`
	DeletedAt       *time.Time `json:"deleted_at_sample,omitempty"`
	IsActive        bool       `json:"is_active_sample"`
}
