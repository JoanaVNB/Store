package database

import(
	"BE-JoanaVidon/user-api/domain"
)

func (r Repository) GetAll (u []domain.User) ([]domain.User, error){
	db := r.DB.Table("user").Find(&u)
	if db.Error != nil{
		return []domain.User{}, db.Error
	}
	return u, nil
}