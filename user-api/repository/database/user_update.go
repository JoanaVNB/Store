package database

import (
	"BE-JoanaVidon/user-api/domain"
)

func (r Repository) Update (id string, phone string, u domain.User) (domain.User, error){
	db := r.DB.Table("user").First(&u, "id = ?", id).Update("phone_number", phone)
	if db.Error != nil{
		return domain.User{}, db.Error
	}
	return u, nil
}