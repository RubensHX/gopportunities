package router

import (
	"github.com/RubensHX/gopportunities/docs"
	"github.com/RubensHX/gopportunities/handler"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	handler.InitializeHandler()
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ReadOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
