package domain

/* type User struct{
	ID string 
	Name string `gorm:"not null" binding:"required"`
	CPF string `gorm:"not null" binding:"required"`
	Email string `gorm:"not null" binding:"required,email"`
	PhoneNumber string `json:"phone_number"`
} */

type User struct{
	ID string `json:"id"`
	Name string `json:"name" binding:"required"`
	CPF string `json:"cpf" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number"`
}