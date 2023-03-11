package service

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"BE-JoanaVidon/user-api/domain"
)

type GetAllUseCase interface {
	Execute([]domain.User) ([]domain.User, error)
}

type GetAllRepository interface {
	GetAll([]domain.User) ([]domain.User, error)
}

type getAllRepository struct{
	getAllRepo GetAllRepository
}

func NewGetAllUseCase (getAllRepo GetAllRepository) *getAllRepository{
	return &getAllRepository{
		getAllRepo: getAllRepo,
	}
}

func (ga getAllRepository) Execute(u []domain.User) ([]domain.User, error){
	users, err := ga.getAllRepo.GetAll(u)
	if err != nil{
		return []domain.User{}, err
	}
	return users, nil
}