package service

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE


type DeleteUseCase interface {
	Execute(string) (error)
}

type DeleteRepository interface {
	Delete(string) (error)
}

type deleteRepository struct{
	deleteRepo DeleteRepository
}

func NewDeleteUseCase (deleteRepo DeleteRepository) *deleteRepository{
	return &deleteRepository{
		deleteRepo: deleteRepo,
	}
}

func (d deleteRepository) Execute(id string) (error){
	err := d.deleteRepo.Delete(id)
	if err != nil{
		return  err
	}
	return nil
}