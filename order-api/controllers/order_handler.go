package controllers

import (
	"BE-JoanaVidon/order-api/domain"
	"BE-JoanaVidon/order-api/presenter"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	domainU "BE-JoanaVidon/user-api/domain"
)

type CreateHandler interface {
	Execute(string, domain.Order, domainU.User) (domain.Order, domainU.User, error)
}

type createHandler struct{
	repository CreateHandler
}

func NewCreateHandler (repository CreateHandler) *createHandler{
	return &createHandler{repository: repository}
}

// Create godoc
// @Summary Create Order
// @Description Create Order
// @ID internal-create-order
// @Tags Internal Order
// @Accept json
// @Produce json
// @Param order body presenter.PresenterOrder true "Create Order"
// @Success 201 {object} presenter.PresenterOrder
// @Failure 400 {object} presenter.GetErrorMsg
// @Failure 409 {object} presenter.GetErrorMsg
// @Failure 412 {object} presenter.GetErrorMsg
// @Failure 500 {object} presenter.GetErrorMsg
// @Router /order [post]
func (cr createHandler) Create(c *gin.Context) {
	var o domain.Order
	var u domainU.User
	var ve validator.ValidationErrors

	if err := c.ShouldBindJSON(&o); err != nil {
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
	userID, _:= c.Params.Get("userID")
	order, user, err := cr.repository.Execute(userID, o, u)
	presenterOrder := *presenter.PresenterOrder(order, user)

	if  err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error to create on domain.Order": err.Error()})
		return
	}
	if order.Item_description == ""{
		c.JSON(http.StatusBadRequest, "Order was not create.")
		return
	}
	c.JSON(http.StatusCreated, presenterOrder)
}
