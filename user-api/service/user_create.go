package service

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"BE-JoanaVidon/user-api/domain"
	"fmt"

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
	if !Validate(u){
		return domain.User{}, fmt.Errorf("some empty field")
	}
	u.ID = uuid.NewString()
	user, err := c.createRepo.Create(u)
	if err != nil{
		return domain.User{}, err
	}
	return user, nil
}