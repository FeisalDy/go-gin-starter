package main

import (
	"log"

	"boiler/config"
	"boiler/internal/database"
	userHandler "boiler/internal/user/handler"
	userRepo "boiler/internal/user/repository"
	userService "boiler/internal/user/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.LoadDBConfig()

	// Initialize database
	database.Init(cfg)

	// Initialize user domain
	userRepository := userRepo.NewUserRepository()
	userService := userService.NewUserService(userRepository)
	userHandler := userHandler.NewUserHandler(userService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// User routes
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUser)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
