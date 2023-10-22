package apis

type Applicant struct {
	ID   uint   `json:"id"`
	Name string `json:"applicant_name"`
}

type Company struct {
	ID   uint   `json:"id"`
	Name string `json:"company_name"`
	Jobs []Job  `json:"jobs"`
}

type Job struct {
	ID          uint   `json:"id"`
	Name        string `json:"job_name"`
	Description string `json:"description"`
	City        string `json:"city"`
	Wage        string `json:"wage"`
	Date        string `json:"date"`
	Company     Company
}
