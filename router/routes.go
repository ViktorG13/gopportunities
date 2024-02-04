package router

import (
	"github.com/ViktorG13/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

/*
*

	Func "initializeRoutes". This Func Contains All Routes.
	@params:
	 * router: pointer of a gin.Engine instance

*
*/
func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{

		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.GET("/openings", handler.ListOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

	}
}
