package api

import (
	"fmt"
	"log"
	"net"

	"github.com/alberthaciverdiyev1/goldenfruit/internal/database"
	"github.com/alberthaciverdiyev1/goldenfruit/internal/http/handler"
	"github.com/alberthaciverdiyev1/goldenfruit/internal/http/middleware"
	"github.com/alberthaciverdiyev1/goldenfruit/internal/services"
	"github.com/gin-gonic/gin"
)

var PortChan = make(chan int, 1)

func Start() {
	db := database.ConnectToDatabase()
	jwtService := services.NewJWTService()
	customerService := services.NewCustomerService(db)
	customerHandler := handler.NewCustomerHandler(customerService)
	authService := services.NewUserService(db, jwtService)
	authHandler := handler.NewUserHandler(authService)

	gin.SetMode(gin.ReleaseMode)
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

	ports := []int{8080, 8081, 8082, 8083, 8084}
	var listener net.Listener
	var err error
	finalPort := 0

	for _, p := range ports {
		addr := fmt.Sprintf("127.0.0.1:%d", p)
		listener, err = net.Listen("tcp", addr)
		if err == nil {
			finalPort = p
			break
		}
		log.Printf("Port %d dolu idi, növbəti yoxlanılır...", p)
	}

	if finalPort == 0 {
		log.Fatal("Heç bir boş port tapılmadı (8080-8084).")
	}

	PortChan <- finalPort

	log.Printf("API Server http://127.0.0.1:%d ünvanında başladıldı", finalPort)

	if err := r.RunListener(listener); err != nil {
		log.Fatalf("Server işə düşə bilmədi: %v", err)
	}
}
