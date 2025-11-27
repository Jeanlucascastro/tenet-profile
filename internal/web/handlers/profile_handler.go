package handlers

import (
	"net/http"
	"strconv"
	"tenet-profile/internal/model"
	service "tenet-profile/internal/services"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	service *service.TenetProfileService
}

func NewProfileHandler(service *service.TenetProfileService) *ProfileHandler {
	return &ProfileHandler{
		service: service,
	}
}

func (h *ProfileHandler) CreateProfile(ctx *gin.Context) {
	var profileDTO model.ProfileDTO

	if err := ctx.ShouldBindJSON(&profileDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := h.service.Save(&profileDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, profile)

}

func (h *ProfileHandler) GetProfileByUserID(ctx *gin.Context) {
	userIDParam := ctx.Param("userId")

	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	profile, err := h.service.GetAllByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, profile)
}
