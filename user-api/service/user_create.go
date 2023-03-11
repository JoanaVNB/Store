package service

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"BE-JoanaVidon/user-api/domain"
	"github.com/google/uuid"
)
type CreateUseCase interface{
	Execute (domain.User) (domain.User, error)
}

type CreateRepository interface {
	Create(domain.User) (domain.User, error)
}

type createRepository struct{
	createRepo CreateRepository
}

func NewCreateUseCase (createRepo CreateRepository) *createRepository{
	return &createRepository{
		createRepo: createRepo,
	}
}

func (c createRepository) Execute(u domain.User) (domain.User, error){
	u.ID = uuid.NewString()
	user, err := c.createRepo.Create(u)
	if err != nil{
		return domain.User{}, err
	}
	user.ID = u.ID
	return user, nil
}