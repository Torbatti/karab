# کاراب
- کاراب یک نمونه از سایت های کاریابی است که میتوان به سبکی جدید ساخت
- کاراب نسخه بهینه شده پروتوتایپ کاراب است که هدف  کارکرد 
زبان go , fiber , sqlite , gorm و htmx را نشان میدهد

# نحوه استفاده
```bash
git clone https://github.com/Torbatti/karab.git
cd karab
go run main.go
```

# models
```
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
	Description  string `json:"description"`
	City         string `json:"city"`
	Wage         string `json:"wage"`
	Date         string `json:"date"`
	CompanyRefer uint   `json:"company_refer"`
}
```
# نمایش
- [فیلم استفاده از کاراب](/showcase/showcase.webm)
- صفحه اصلی
![](/showcase/SS1.png)
- صفحه ورود
![](/showcase/SS2.png)
- صفحه جست و جوی فرصت های شغلی
![](/showcase/SS3.png)
- صفحه ی فرصت شغلی انتخاب شده
![](/showcase/SS4.png)




