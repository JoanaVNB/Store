package inMemo

import (
	"BE-JoanaVidon/user-api/domain"
	"errors"
)

func (in InMemoShopRepository) Update(id string, phone string, u domain.User) (domain.User, error){
	if _, ok := in.uMap[id]; ok {
		in.uMap[id] = u
		u.PhoneNumber = phone
	} else {
		return domain.User{}, errors.New("not found")
	}
	return u, nil
}


