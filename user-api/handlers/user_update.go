package handlers

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/presenter"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

type UpdateHandler interface {
	Execute(string, string, domain.User) (domain.User, error)
}

type updateHandler struct{
	repository UpdateHandler
}

func NewUpdateHandler (repository UpdateHandler) *updateHandler{
	return &updateHandler{repository: repository}
}

// Create godoc
// @Summary Update User
// @Description Update User
// @ID internal-Update-user
// @Tags Internal User
// @Accept json
// @Produce json
// @Param user body presenter.PresenterUser true "Presenter User"
// @Success 201 {object} presenter.PresenterUser
// @Failure 400 {object} presenter.GetErrorMsg
// @Failure 409 {object} presenter.GetErrorMsg
// @Failure 412 {object} presenter.GetErrorMsg
// @Failure 500 {object} presenter.GetErrorMsg
// @Router /users/:id [put]
func (up updateHandler) Update(c *gin.Context) {
	var u domain.User
	var ve validator.ValidationErrors
	givenID := c.Params.ByName("id")
	givenPhone := c.Params.ByName("phone")

	if err := c.ShouldBindJSON(&u); err != nil {
		if errors.As(err, &ve){
			out := make([]presenter.ErrorMsg, len(ve))
			for i, fe := range ve{
				out[i] = presenter.ErrorMsg{Field: fe.Field(), Message: presenter.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": out})
		}
		return 
	}

	user, err := up.repository.Execute(givenID, givenPhone, u)
	if err != nil{
		c.JSON(http.StatusBadRequest, "não encontrado")
		return
	}
	presenterUser := *presenter.PresenterUser(user)
	c.JSON(http.StatusOK, presenterUser)
}
