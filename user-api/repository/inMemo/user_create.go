package inMemo

import (
	"BE-JoanaVidon/user-api/domain"
)

func (in InMemoRepository) Create(u domain.User) (domain.User, error){
	in.uMap[u.ID] = u
	return u,  nil
}
