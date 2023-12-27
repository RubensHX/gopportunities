package handler

import (
	"fmt"
	"net/http"

	"github.com/RubensHX/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.ErrorF("Error finding opening: %v", err)
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with id %s not found", id))
		return
	}

	if err := db.Delete(&schemas.Opening{}, id).Error; err != nil {
		logger.ErrorF("Error deleting opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error deleting opening with id %s", id))
		return
	}

	sendSuccess(ctx, opening)

}
