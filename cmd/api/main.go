// @title Go Clean Architecture API
// @version 1.0
// @description This is a sample server for a Go application following Clean Architecture principles.
// @host localhost:8080
// @BasePath /api/v1

package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/Bobby-P-dev/go-clean-api.git/docs"
	"github.com/Bobby-P-dev/go-clean-api.git/internal/user"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	{
		// api.GET("/health", func(c *gin.Context) {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"sucsses": true,
		// 		"message": "API is running",
		// 	})
		// })

		userHandler := user.NewHandler()
		api.POST("/users", userHandler.CreateUser)
	}
	r.Run(":8080")
}
