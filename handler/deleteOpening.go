package handler

import (
	"net/http"

	"github.com/RubensHX/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "Invalid id")
		return
	}

	var opening schemas.Opening
	if err := db.First(&opening, id).Error; err != nil {
		logger.ErrorF("Error finding opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Error finding opening")
		return
	}

	if err := db.Delete(&schemas.Opening{}, id).Error; err != nil {
		logger.ErrorF("Error deleting opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Error deleting opening")
		return
	}

	sendSuccess(ctx, nil)


}
