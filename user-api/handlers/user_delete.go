package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
type DeleteHandler interface {
	Execute(string) (error)
}

type deleteHandler struct{
	repository DeleteHandler
}

func NewDeleteHandler (repository DeleteHandler) *deleteHandler{
	return &deleteHandler{repository: repository}
}

// Delete godoc
// @Summary Delete User
// @Description Delete User
// @ID internal-delete-user
// @Tags Internal User
// @Success 201 {object} 
// @Failure 400 {object}
// @Failure 409 {object}
// @Failure 412 {object}
// @Failure 500 {object}
// @Router /users/:id [delete]
func (d deleteHandler) Delete(c *gin.Context) {
	givenID := c.Params.ByName("id")
	err := d.repository.Execute(givenID)
	if err != nil{
		c.JSON(http.StatusBadRequest, "not found")
		return
	}
	c.JSON(http.StatusOK, "deleted")
}