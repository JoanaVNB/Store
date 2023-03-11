package presenter

import (
	"BE-JoanaVidon/order-api/domain"
	domainU "BE-JoanaVidon/user-api/domain"
)

type PresentOrder struct{
	ID string 
	User_id string `json:"user_id"`
	Item_description string `json:"item_description"`
	Item_quantity int `json:"item_quantity"`
	Item_price int `json:"item_price"`
	Total_value int `json:"total_value"`
	NameUser string `json:"name_user"`
	CPFUser string `json:"cpf_user"`
	EmailUser string `json:"email_user"`
	PhoneNumberUser string `json:"phone_user"`
}

func PresenterOrder(o domain.Order, u domainU.User) *PresentOrder{
	return &PresentOrder{
		ID:               o.ID,
		User_id:          u.ID,
		Item_description: o.Item_description,
		Item_quantity:    o.Item_quantity,
		Item_price:       o.Item_price,
		Total_value:      o.Total_value,
		NameUser:         u.Name,
		CPFUser:          u.CPF,
		EmailUser:        u.Email,
		PhoneNumberUser:  u.PhoneNumber,
	}
}