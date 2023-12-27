package handler

import (
	"net/http"

	"github.com/RubensHX/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Error validating request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusNotFound, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		logger.ErrorF("Error finding opening: %v", err)
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.ErrorF("Error updating opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Error updating opening")
		return
	}

	sendSuccess(ctx, opening)
}


