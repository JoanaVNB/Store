package usecase

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"BE-JoanaVidon/order-api/domain"
	domainU"BE-JoanaVidon/user-api/domain"
	"github.com/google/uuid"
)

type CreateUseCase interface{
	Execute(string, domain.Order, domainU.User) (domain.Order, domainU.User, error)
}

type CreateRepository interface {
	Create(string, domain.Order, domainU.User) (domain.Order, domainU.User, error)
}

type createRepository struct{
	createRepo CreateRepository
}

func NewCreateUseCase (createRepo CreateRepository) *createRepository{
	return &createRepository{
		createRepo: createRepo,
	}
}

func (c createRepository) Execute(id string, o domain.Order, u domainU.User) (domain.Order, domainU.User, error){
	o.ID = uuid.NewString()
	order, user, err := c.createRepo.Create(id, o, u)
	if err != nil{
		return domain.Order{}, user, err
	}
	order.ID =o.ID
	return order, user, nil
}