package inMemo

import (
	"BE-JoanaVidon/user-api/domain"
)

func (in InMemoShopRepository) GetAll([]domain.User) ([]domain.User, error){
	var users []domain.User
	for _, k := range in.uMap{
		user := k
		users = append(users, user)
	}	
	return users, nil
}
