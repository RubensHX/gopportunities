package handler

import (
	"net/http"

	"github.com/RubensHX/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		logger.ErrorF("Error finding openings: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Error finding openings")
		return
	}

	sendSuccess(ctx, openings)
}
