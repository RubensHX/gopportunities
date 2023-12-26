package handler

import (
	"net/http"

	"github.com/RubensHX/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("Error validating request: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create((&opening)).Error; err != nil {
		logger.ErrorF("Error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating opening in database")
		return
	}

	sendSuccess(ctx, opening)
}
