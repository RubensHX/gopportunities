package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := db.Create((&request)).Error; err != nil {
		logger.ErrorF("Error creating opening: %v", err)
		ctx.JSON(500, gin.H{
			"message": "Error creating opening",
		})
		return
	}
}
