package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
*

	Func "initializeRoutes".
	@params:
	 * router: pointer of a gin.Engine instance

*
*/
func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{

		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello",
			})
		})

		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello",
			})
		})

		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello",
			})
		})

		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello",
			})
		})

		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Hello",
			})
		})

	}
}
