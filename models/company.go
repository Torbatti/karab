package models

import "time"

type Company struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time
	CompanyName string `json:"company_name"`
}
