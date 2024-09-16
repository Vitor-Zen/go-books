package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o Router utilizando as configurações Default do gin
	router := gin.Default()
	// Define uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Inicia o servidor
	router.Run(":8080")
}
