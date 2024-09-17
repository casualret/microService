package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"urlshortener/internal/models"
)

func (h *Handlers) CreateTag(c *gin.Context) {
	const op = "handlers.CreateTag"

	var input models.Tag
	err := c.BindJSON(&input)
	if err != nil {
		h.Logger.Error("Failed to bind JSON: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	//ctx, cancel := context.WithTimeout(context.Background(), contextTimeResponse)
	//defer cancel()

	err = h.App.CreateTag(input)
	if err != nil {
		h.Logger.Error("Error creating tag: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusInternalServerError, internalServerError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) CreateFeature(c *gin.Context) {
	const op = "handlers.CreateFeature"

	var input models.Feature
	err := c.BindJSON(&input)
	if err != nil {
		h.Logger.Error("Failed to bind JSON: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	err = h.App.CreateFeature(input)
	if err != nil {
		h.Logger.Error("Error creating feature: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusInternalServerError, internalServerError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
