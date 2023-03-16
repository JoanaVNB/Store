package domain

import(
	"time"
)

type User struct{
	ID string `json:"id" gorm:"varchar(150); not null"`
	Name string `json:"name" binding:"required" gorm:"varchar(150); not null"`
	CPF string `json:"cpf" binding:"required" gorm:"varchar(150); not null; unique"`
	Email string `json:"email" binding:"required,email" gorm:"varchar(150); not null; unique"`
	PhoneNumber string `json:"phone_number" gorm:"varchar(18); not null; unique"`
	CreatedAt time.Time `json:"created_at" gorm:"default current_timestamp()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default current_timestamp()"`
}

func (User) TableName() string{
	return "user"
}