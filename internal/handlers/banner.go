package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"urlshortener/internal/models"
)

func (h *Handlers) CreateBanner(c *gin.Context) { // Добавить связь many-to-many
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

func (h *Handlers) GetBanners(c *gin.Context) {
	const op = "handlers.GetBanners"

	var req models.GetBannersReq

	featureIDStr := c.Query("feature_id")
	tagIDStr := c.Query("tag_id")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	//fmt.Println("featureIDSTR " + featureIDStr)
	//fmt.Println("tagIDSTR " + tagIDStr)
	//fmt.Println("limitIDSTR " + limitStr)
	//fmt.Println("offserIDSTR " + offsetStr)

	if featureIDStr != "" {
		featureID, err := strconv.Atoi(featureIDStr)
		if err != nil {
			h.Logger.Error("Failed get feature_id:", fmt.Errorf("%s: %w", op, err))
			newErrorResponse(c, http.StatusBadRequest, incorrectData)
			return
		}
		featureIDPtr := &featureID
		req.FeatureID = featureIDPtr

	}

	if tagIDStr != "" {
		tagID, err := strconv.Atoi(tagIDStr)
		if err != nil {
			h.Logger.Error("Failed get tag_id:", fmt.Errorf("%s: %w", op, err))
			newErrorResponse(c, http.StatusBadRequest, incorrectData)
			return
		}
		tagIDPtr := &tagID
		req.TagID = tagIDPtr
	}

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			h.Logger.Error("Failed get limit: ", fmt.Errorf("%s: %v", op, err))
			newErrorResponse(c, http.StatusBadRequest, incorrectData)
			return
		}
		limitPtr := &limit
		req.Limit = limitPtr
	}

	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			h.Logger.Error("Failed get offset: ", fmt.Errorf("%s: %v", op, err))
			newErrorResponse(c, http.StatusBadRequest, incorrectData)
			return
		}
		offsetPtr := &offset
		req.Offset = offsetPtr
	}

	banners, err := h.App.GetBanners(req)
	if err != nil {
		h.Logger.Error("Error getting banners: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	//c.String(200, "Пока я мастерил фрегат мир стал бессмыслено богат и полон гнуси")
	c.JSON(http.StatusOK, banners)
}
