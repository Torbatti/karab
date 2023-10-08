package models

import "time"

type Applicant struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time
	ApplicantName string `json:"applicant_name"`
}
