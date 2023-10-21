package models

import "time"

type Applicant struct {
	ID            uint `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time
	ApplicantName string `json:"applicant_name"`
}

type Company struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name      string `json:"company_name"`
	Jobs      []Job  `gorm:"foreignKey:CompanyRefer"`
}

type Job struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Name         string `json:"job_name"`
	Description  string `json:"Description"`
	CompanyRefer uint
}
