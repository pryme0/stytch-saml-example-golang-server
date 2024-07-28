package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	routes "saml_sso/internal/controllers"
	"saml_sso/internal/database"
)

const PORT = ":3002"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, datbaseErr := database.Connect()

	if datbaseErr != nil {
		// Handle error
		log.Fatal("Something went wrong connecting to database")
	}

	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},                                                 // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},                      // Allow specific methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "authorization"}, // Allow specific headers
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	r.Use(cors.New(corsConfig))

	apiV1 := r.Group("/api/v1")
	{
		routes.TenantRoutes(apiV1, db)

		routes.MemberRoutes(apiV1, db)

		routes.AuthRoutes(apiV1, db)
	}

	fmt.Println("Server is listening on port", PORT)
	r.Run(PORT)
}
