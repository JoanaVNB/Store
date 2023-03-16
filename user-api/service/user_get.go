package service

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"BE-JoanaVidon/user-api/domain"
)

type GetUseCase interface {
	Execute(string, domain.User) (domain.User, error)
}

type GetRepository interface {
	Get(string, domain.User) (domain.User, error)
}

type getRepository struct{
	getRepo GetRepository
}

func NewGetUseCase (getRepo GetRepository) *getRepository{
	return &getRepository{
		getRepo: getRepo,
	}
}

func (g getRepository) Execute(id string, u domain.User) (domain.User, error){
	user, err := g.getRepo.Get(id, u)
	if err != nil{
		return domain.User{}, err
	}
	user.ID = u.ID
	return user, nil
}