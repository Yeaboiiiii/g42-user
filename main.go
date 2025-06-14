package main

import (
	"log"
	"time"

	"g42-user/cmd/handler"
	"g42-user/cmd/logic"
	"g42-user/repositories"
	"g42-user/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB connection
	userRepo, err := repositories.NewUserRepository("mongodb://localhost:27017", "g42-user", "auth")
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Initialize business logic
	//	userRepo.CreateUser("test@test.com", "123456")

	userLogic := logic.NewUserLogic(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userLogic)

	// Setup Gin router
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Public routes
	router.POST("/signup", userHandler.Signup)
	router.POST("/login", userHandler.Login)
	router.GET("/logout", userHandler.Logout)

	// Protected routes
	protected := router.Group("/")
	protected.Use(utils.AuthMiddleware())
	{
		protected.POST("/user/details", userHandler.GetUserDetails)
		protected.GET("/user", userHandler.GetUserDetailsByID)
	}

	// Start server
	if err := router.Run(":8002"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
