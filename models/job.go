package models

import "time"

type List struct {
	ID             uint `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time
	CompanyRefer   int       `json:"company_id"`
	Company        Company   `gorm:"foreignKey:CompanyRefer"`
	ApplicantRefer int       `json:"applicant_id"`
	Applicant      Applicant `gorm:"foreignKey:ApplicantRefer"`
}
