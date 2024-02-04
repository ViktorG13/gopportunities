package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Iniciando o Router.
	router := gin.Default()

	// Definindo a Rota GET->"/ping".
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Executando o Router.
	router.Run(":8080")
}
