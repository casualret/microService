package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"urlshortener/internal/models"
)

func (h *Handlers) CreateBanner(c *gin.Context) {
	const op = "handlers.CreateBanner"

	var input models.CreateBannerReq

	err := c.BindJSON(&input)
	if err != nil {
		h.Logger.Error("Failed to bind JSON: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	//Проверка Feature
	_, err = strconv.Atoi(input.FeatureID)
	if err != nil {
		h.Logger.Error("Incorrect feature_id: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	//Проверка Active
	if !strings.EqualFold(input.IsActive, "true") && !strings.EqualFold(input.IsActive, "false") {
		h.Logger.Error("Incorrect is_active: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	// Проверка массива тегов

	for _, tag := range input.TagIds {
		if _, err := strconv.Atoi(tag); err != nil {
			h.Logger.Error("Invalid tag value at index", fmt.Errorf("%s: %v", op, err))
			newErrorResponse(c, http.StatusBadRequest, incorrectData)
			return
		}
	}

	err = h.App.CreateBanner(input)
	if err != nil {
		h.Logger.Error("Error creating banner: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusInternalServerError, internalServerError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
