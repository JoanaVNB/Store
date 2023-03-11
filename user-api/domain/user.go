package domain

type User struct{
	ID string 
	Name string `gorm:"not null" binding:"required"`
	CPF string `gorm:"not null" binding:"required"`
	Email string `gorm:"not null" binding:"required,email"`
	PhoneNumber string `json:"phone_number"`
}
