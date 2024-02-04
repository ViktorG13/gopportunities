package router

import (
	docs "github.com/ViktorG13/gopportunities/docs"
	"github.com/ViktorG13/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

/*
*

	Func "initializeRoutes". This Func Contains All Routes.
	@params:
	 * router: pointer of a gin.Engine instance

*
*/
func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializerHandle()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{

		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.GET("/openings", handler.ListOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
