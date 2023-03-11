package service

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"BE-JoanaVidon/user-api/domain"
)

type UpdateUseCase interface {
	Execute(string, string, domain.User) (domain.User, error)
}

type UpdateRepository interface {
	Update(string, string, domain.User) (domain.User, error)
}

type updateRepository struct{
	updateRepo UpdateRepository
}

func NewUpdateUseCase (updateRepo UpdateRepository) *updateRepository{
	return &updateRepository{
		updateRepo: 	updateRepo,
	}
}

func (up updateRepository) Execute(id string, phone string, u domain.User ) (domain.User, error){
	user, err := up.updateRepo.Update(id, phone, u)
	if err != nil{
		return domain.User{}, err
	}
	return user, nil
}