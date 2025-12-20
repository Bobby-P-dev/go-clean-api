// @title Go Clean Architecture API
// @version 1.0
// @description This is a sample server for a Go application following Clean Architecture principles.
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /api/v1

package main

import (
	"fmt"
	"github.com/Bobby-P-dev/go-clean-api/internal/article"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/Bobby-P-dev/go-clean-api/docs"
	"github.com/Bobby-P-dev/go-clean-api/internal/config"
	"github.com/Bobby-P-dev/go-clean-api/internal/user"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found, using environment variables")
	}
	cfg := config.LoadConfig()
	db := config.InitDB(cfg)

	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		panic("failed to connect database: " + err.Error())
	}
	fmt.Println("Database connected successfully")

	if err := db.AutoMigrate(&user.User{}, &article.Article{}); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	articleRepo := article.NewRepository(db)
	articleService := article.NewService(articleRepo)
	articleHandler := article.NewHandler(articleService)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	{
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users", userHandler.ListUsers)
		api.POST("/login", userHandler.LoginUser)
	}

	articleGroup := r.Group("/article")
	articleGroup.Use(user.AuthMiddleware())
	{
		articleGroup.POST("/create", articleHandler.CreateArticle)
	}

	protected := r.Group("/protected/api/v1")
	protected.Use(user.AuthMiddleware())
	{
		protected.GET("/users/me", userHandler.GetMe)
	}
	if err := r.Run(os.Getenv("APP_PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
