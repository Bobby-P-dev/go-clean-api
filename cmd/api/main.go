// @title Go Clean Architecture API
// @version 1.0
// @description This is a sample server for a Go application following Clean Architecture principles.
// @host localhost:8080
// @BasePath /api/v1

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	_ "github.com/Bobby-P-dev/go-clean-api.git/docs"
	"github.com/Bobby-P-dev/go-clean-api.git/internal/config"
	"github.com/Bobby-P-dev/go-clean-api.git/internal/user"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	cfg := config.LoadConfig()
	db := config.InitDB(cfg)

	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		panic("failed to connect database: " + err.Error())
	}
	fmt.Println("Database connected successfully")

	if err := db.AutoMigrate(&user.User{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

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
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users", userHandler.ListUsers)
		api.POST("/login", userHandler.LoginUser)
	}

	r.Run(":8080")
}
