package handler

import (
	"github.com/RubensHX/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func sendSuccess(ctx *gin.Context, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SuccessResponse struct {
	Code int                     `json:"code"`
	Data schemas.OpeningResponse `json:"data"`
}
