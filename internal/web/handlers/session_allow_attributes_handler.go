package handlers

import (
	"net/http"
	"strconv"
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

func (h *SessionAllowAttributesHandler) UpdateSessionAllowAttributes(ctx *gin.Context) {
	var saa model.SessionAllowAttributes

	if err := ctx.ShouldBindJSON(&saa); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sessionIdParam := ctx.Param("sessionId")

	id, err := strconv.ParseUint(sessionIdParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
		return
	}

	saa.SessionID = int64(id)

	updateError := h.service.Update(&saa, int64(id))
	if updateError != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": updateError.Error()})
		return
	}

	ctx.JSON(http.StatusOK, saa)
}

func (h *SessionAllowAttributesHandler) findSessionAllowAttributesByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	saa, getError := h.service.GetByID(id)
	if getError != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": getError.Error()})
		return
	}

	ctx.JSON(http.StatusOK, saa)
}

func (h *SessionAllowAttributesHandler) GetSessionAllowAttributesBySessionIdAndUserId(ctx *gin.Context) {

	sessionIdParam := ctx.Param("sessionId")
	userIdParam := ctx.Param("userId")

	sessionID, err := strconv.ParseUint(sessionIdParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
		return
	}

	userID, err := strconv.ParseUint(userIdParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	saa, getError := h.service.FindBySessionIdAndUserWithThisAttribute(int64(sessionID), int64(userID))
	if getError != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": getError.Error()})
		return
	}

	ctx.JSON(http.StatusOK, saa)
}
