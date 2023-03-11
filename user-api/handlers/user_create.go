package handlers

import (
	"BE-JoanaVidon/user-api/domain"
	"BE-JoanaVidon/user-api/presenter"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

type CreateHandler interface {
	Execute(domain.User) (domain.User, error)
}

type createHandler struct{
	repository CreateHandler
}

func NewCreateHandler (repository CreateHandler) *createHandler{
	return &createHandler{repository: repository}
}

// Create godoc
// @Summary Create User
// @Description Create User
// @ID internal-create-user
// @Tags Internal User
// @Accept json
// @Produce json
// @Param user body presenter.PresenterUser true "Create User"
// @Success 201 {object} presenter.PresenterUser
// @Failure 400 {object} presenter.GetErrorMsg
// @Failure 409 {object} presenter.GetErrorMsg
// @Failure 412 {object} presenter.GetErrorMsg
// @Failure 500 {object} presenter.GetErrorMsg
// @Router /users [post]
func (cr createHandler) Create(c *gin.Context) {
	var u domain.User
	var ve validator.ValidationErrors

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
	user, err := 	cr.repository.Execute(u)
	presenterUser := *presenter.PresenterUser(user)
	if  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error to create on domain.User": err.Error()})
		return
	}
	if user.Name == ""{
		c.JSON(http.StatusBadRequest, "User was not create.")
		return
	}
	c.JSON(http.StatusCreated, presenterUser)
}