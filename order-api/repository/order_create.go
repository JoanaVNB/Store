package repository

import (
	"BE-JoanaVidon/order-api/domain"
	domainU "BE-JoanaVidon/user-api/domain"
	"fmt"

	"gorm.io/gorm"
)

type Driver = *gorm.DB

type Repository struct{
	DB Driver
	err error
}

func NewRepository (DB Driver, err error) *Repository{
	return &Repository{DB: DB, err: err}
}

func (r Repository) Create(id string, o domain.Order, u domainU.User) (domain.Order, domainU.User, error){
	db := r.DB.Table("order").Create(&o)
	if db.Error != nil{
		return domain.Order{}, u, db.Error
	}
	dbUser := r.DB.Table("user").First(&u, "id = ?", id)
	if db.Error != nil{
		fmt.Println("error to find user")
		return domain.Order{}, domainU.User{}, dbUser.Error
	}
	return o, u, nil
}


