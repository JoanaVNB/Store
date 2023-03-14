package handlers

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/presenter"
	"net/http"
	"github.com/gin-gonic/gin"
)

type GetAllHandler interface {
	Execute([]domain.User) ([]domain.User, error)
}

type getAllHandler struct{
	repository GetAllHandler
}

func NewGetAllHandler (repository GetAllHandler) *getAllHandler{
	return &getAllHandler{repository: repository}
}

// GetAll godoc
// @Summary GetAll User
// @Description GetAll User
// @ID internal-GetAll-user
// @Tags Internal User
// @Accept json
// @Produce json
// @Param user body presenter.PresenterUser true "GetAll User"
// @Success 201 {object} presenter.PresenterUser
// @Failure 400 {object} presenter.GetErrorMsg
// @Failure 409 {object} presenter.GetErrorMsg
// @Failure 412 {object} presenter.GetErrorMsg
// @Failure 500 {object} presenter.GetErrorMsg
// @Router /users [get]
func (ga getAllHandler) GetAll(c *gin.Context){
	var u []domain.User

	users, err := ga.repository.Execute(u)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		return
	}

	for _, eachUser := range users{
		user := eachUser
		presenterUser := *presenter.PresenterUser(user)
		c.JSON(http.StatusOK, presenterUser)
	}
}
