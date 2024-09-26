package handlers

import (
	"errors"
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

	//Считываем аргументы для запроса

	featureIDStr := c.Query("feature_id")
	tagIDStr := c.Query("tag_id")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	// Проверяем наличие Feature, создаем переменную, передаем указатель на нее в структуру

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

	// Для tag

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

	// Для Limit

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

	//Для offset

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

	// Вызываем функцию сервиса для получения баннеров

	banners, err := h.App.GetBanners(req)
	if err != nil {
		h.Logger.Error("Error getting banners: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	//c.String(200, "Пока я мастерил фрегат мир стал бессмыслено богат и полон гнуси")

	// Отправляем массив структур в body

	c.JSON(http.StatusOK, banners)
}

func (h *Handlers) DeleteBanner(c *gin.Context) {
	const op = "handlers.DeleteBanner"

	// TODO: FIX PARAM CHECK
	bannerIDStr := c.Param("id")
	if bannerIDStr == "" {
		h.Logger.Error("Incorrect data: ", fmt.Errorf("%s: %v", op, errors.New("ID parameter is missing")))
		newErrorResponse(c, http.StatusBadRequest, bannerNotSelected)
		return
	}

	bannerID, err := strconv.ParseInt(bannerIDStr, 10, 64)
	if err != nil {
		h.Logger.Error("Failed convert bannerID to int64: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	err = h.App.DeleteBanner(bannerID)
	if err != nil {
		h.Logger.Error("Error deleting banner: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}
	c.JSON(http.StatusAccepted, nil)
}

func (h *Handlers) PatchBanner(c *gin.Context) {
	const op = "handlers.PatchBanner"
}
