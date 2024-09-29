package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"microService/internal/models"
	"net/http"
)

func (h *Handlers) SignUp(c *gin.Context) {
	const op = "handlers.SignUp"

	var input models.CreateUserReq
	err := c.BindJSON(&input)
	if err != nil {
		h.Logger.Error("Failed to bind JSON: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	err = h.App.SignUp(input)
	if err != nil {
		h.Logger.Error("Error creating User: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusInternalServerError, internalServerError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) SignIn(c *gin.Context) {

}
