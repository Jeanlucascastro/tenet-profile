package handlers

import (
	"net/http"
	"tenet-profile/internal/model"
	service "tenet-profile/internal/services"

	"github.com/gin-gonic/gin"
)

type SessionAllowAttributesHandler struct {
	service *service.TenetSessionAllowAttributesService
}

func NewSessionAllowAttributesHandler(service *service.TenetSessionAllowAttributesService) *SessionAllowAttributesHandler {
	return &SessionAllowAttributesHandler{
		service: service,
	}
}

func (h *SessionAllowAttributesHandler) CreateSessionAllowAttributes(ctx *gin.Context) {
	var saa model.SessionAllowAttributes

	if err := ctx.ShouldBindJSON(&saa); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Save(&saa)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, saa)
}
