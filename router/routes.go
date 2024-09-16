package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/books", func(ctx *gin.Context) {
			// Handle GET request to /api/v1/books
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Books",
			})
		})
	}
}
