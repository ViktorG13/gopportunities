package router

import "github.com/gin-gonic/gin"

/*
*

	Func "Initialize". This Func Initialize The Router.

*
*/
func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Run The Server.
	router.Run(":8080")
}
