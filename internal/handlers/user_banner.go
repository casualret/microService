package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"microService/internal/models"
	"net/http"
)

func (h *Handlers) GetUserBanner(c *gin.Context) {
	const op = "handlers.GetUserBanner"

	featureId, ok := c.GetQuery("feature_id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}
	tagId, ok := c.GetQuery("tag_id")
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}
	req := models.GetUserBannerReq{FeatureID: featureId, TagID: tagId}

	banner, err := h.App.GetUserBanner(req)
	if err != nil {
		h.Logger.Error("Error getting banner: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	c.JSON(http.StatusOK, banner)
}
