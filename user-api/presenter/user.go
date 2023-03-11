package presenter

import (
	"BE-JoanaVidon/user-api/domain"
)

type PresentUser struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	CPF string `json:"cpf,omitempty"`
	Email string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
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