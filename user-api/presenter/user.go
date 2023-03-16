package presenter

import (
	"BE-JoanaVidon/user-api/domain"
	"time"
)

type PresentUser struct {
	ID string `json:"id"`
	Name string `json:"name"`
	CPF string `json:"cpf"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type GetUser struct {
	ID string `json:"id"`
	Name string `json:"name"`
	CPF string `json:"cpf"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"` 
}


func PresenterUser(u domain.User) *PresentUser{
	return &PresentUser{
		ID:          u.ID,
		Name:       u.Name,
		CPF:         u.CPF,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
	}
}