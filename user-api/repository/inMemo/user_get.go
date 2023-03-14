package inMemo

import (
	"BE-JoanaVidon/user-api/domain"
)

func (in InMemoRepository) Get(id string, u domain.User) (domain.User, error){
	for _, k := range in.uMap{
		gotID := k.ID
		if gotID == id{
			return u, nil
		}
	}
	return u, nil
}
