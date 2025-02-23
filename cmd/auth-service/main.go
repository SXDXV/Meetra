package main

import (
	"fmt"
	"log"

	"github.com/SXDXV/meetra/internal/auth/handlers"
	"github.com/SXDXV/meetra/internal/auth/models"
	"github.com/SXDXV/meetra/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Auth Service запущен на порту 8080")

	router := gin.Default()

	router.POST("/auth/register", handlers.RegisterUSer)

	router.Run(":8080")
}
