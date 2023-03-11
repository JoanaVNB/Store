package database

import(
	"BE-JoanaVidon/user-api/domain"
)

func (r Repository) Create(u domain.User) (domain.User, error){
	db := r.DB.Table("user").Create(&u)
	if db.Error != nil{
		return domain.User{}, db.Error
	}
	return u, nil
}