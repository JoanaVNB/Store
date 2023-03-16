package inMemo

import (
	"BE-JoanaVidon/user-api/domain"
)

type InMemoRepository struct{
	uMap	map[string]domain.User
}

func NewInMemoRepository() *InMemoRepository{
	return &InMemoRepository{
		uMap: make(map[string]domain.User)}
}
