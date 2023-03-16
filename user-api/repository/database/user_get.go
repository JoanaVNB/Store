package database

import(
	"BE-JoanaVidon/user-api/domain"
)

func (r Repository) Get(id string, u domain.User) (domain.User, error){
	db := r.DB.Table("user").First(&u, "id = ?", id)
	if db.Error != nil{
		return domain.User{}, db.Error
	}
	return u, nil
}
