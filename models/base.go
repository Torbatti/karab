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
	Name      string `json:"company_name" gorm:"default:شرکت معتبر"`
	Jobs      []Job  `gorm:"foreignKey:CompanyRefer"`
}

type Job struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Name         string `json:"job_name"`
	Description  string `json:"description"`
	City         string `json:"city" gorm:"default:تمامی شهرها"`
	Wage         string `json:"wage" gorm:"default:۱۰-۲۰ میلیون تومان"`
	Date         string `json:"date" gorm:"default:۵ روز پیش"`
	CompanyRefer uint   `json:"company_refer" gorm:"default:2"`
}
