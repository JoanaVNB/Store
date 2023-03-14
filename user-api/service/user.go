package service

type UserUC struct {
	createUC CreateUseCase
	getUC GetUseCase
	getAllUC GetAllUseCase
	updateUC UpdateUseCase
	deleteUC DeleteUseCase
}

func NewUserUC (
	createUC CreateUseCase,
	getUC GetUseCase,
	getAllUC GetAllUseCase,
	updateUC UpdateUseCase,
	deleteUC DeleteUseCase) *UserUC{
		return &UserUC{
			createUC: createUC,
			getUC:    getUC,
			getAllUC: getAllUC,
			updateUC: updateUC,
			deleteUC: deleteUC,
		}
	}
