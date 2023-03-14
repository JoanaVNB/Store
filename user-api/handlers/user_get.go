package handlers

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/presenter"
	"net/http"
	"github.com/gin-gonic/gin"
)

type GetHandler interface {
	Execute(string, domain.User) (domain.User, error)
}

type getHandler struct{
	repository GetHandler
}

func NewGetHandler (repository GetHandler) *getHandler{
	return &getHandler{repository: repository}
}

// Get godoc
// @Summary Get User
// @Description Get User
// @ID internal-get-user
// @Tags Internal User
// @Accept json
// @Produce json
// @Param user body presenter.PresenterUser true "Get User"
// @Success 201 {object} presenter.PresenterUser
// @Failure 400 {object} presenter.GetErrorMsg
// @Failure 409 {object} presenter.GetErrorMsg
// @Failure 412 {object} presenter.GetErrorMsg
// @Failure 500 {object} presenter.GetErrorMsg
// @Router /users/:id [get]
func (g getHandler) Get(c *gin.Context) {
	var u domain.User

	givenID, _:= c.Params.Get("id")
	user, err := g.repository.Execute(givenID, u)
	if err != nil{
		c.JSON(http.StatusBadRequest, "not found")
	}
	presenterUser := *presenter.PresenterUser(user)
	c.JSON(http.StatusOK, presenterUser)
}