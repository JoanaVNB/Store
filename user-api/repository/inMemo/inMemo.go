package inMemo

import (
	"BE-JoanaVidon/user-api/domain"
)

type InMemoShopRepository struct{
	uMap	map[string]domain.User
}

func NewInMemoShopRepository() *InMemoShopRepository{
	return &InMemoShopRepository{uMap: make(map[string]domain.User)}
}

