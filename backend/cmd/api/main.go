package main

import (
	"log"

	"github.com/alberthaciverdiyev1/goldenfruit/internal/database"
	"github.com/alberthaciverdiyev1/goldenfruit/internal/http/handler"
	"github.com/alberthaciverdiyev1/goldenfruit/internal/http/middleware"
	"github.com/alberthaciverdiyev1/goldenfruit/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectToDatabase()

	jwtService := services.NewJWTService()

	customerService := services.NewCustomerService(db)
	customerHandler := handler.NewCustomerHandler(customerService)

	authService := services.NewUserService(db, jwtService)
	authHandler := handler.NewUserHandler(authService)

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/logout", authHandler.Logout)
		}

		customers := api.Group("/customers")
		customers.Use(middleware.AuthMiddleware(jwtService))
		{
			customers.GET("/", customerHandler.List)
			customers.GET("/:id", customerHandler.GetByID)
			customers.POST("/", customerHandler.Create)
			customers.PUT("/:id", customerHandler.Update)
			customers.DELETE("/:id", customerHandler.Delete)
		}
	}

	log.Println("Server 8080 portunda başlatıldı...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Sunucu başlatılamadı: %v", err)
	}
}
