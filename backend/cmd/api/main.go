package main

import (
	"ecommerce/internal/adapter/auth"
	repo "ecommerce/internal/adapter/db"
	adapter "ecommerce/internal/adapter/http"
	"ecommerce/internal/core/product"
	"ecommerce/internal/core/user"
	"ecommerce/pkg/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("🚀 Iniciando servidor...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.GetConnection()

	r := gin.New()
	r.Use(gin.Recovery()) // tratamento de panic
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type"},
	}))

	productRepo := repo.NewPostgresRepo(db)
	service := product.NewService(productRepo)
	productHandler := adapter.NewHandler(service)
	productHandler.Register(r)

	userRepo := repo.NewUserRepo(db)
	userService := user.NewService(userRepo, auth.NewAuth("aspdaskdpasd"))
	userHandler := adapter.NewUserHandler(userService)
	userHandler.RegisterRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic("Erro ao rodar o servidor: " + err.Error())
	}
}
