package database

import(
	"BE-JoanaVidon/user-api/domain"
)

func (r Repository) Delete (id string) (error){
	var u domain.User
	db := r.DB.Table("user").First(&u, "id = ?", id).Delete(&u)
	if db.Error != nil{
		return db.Error
	}
	return nil
}