package main

import (
	"net/http"
	"sample-api-go/internal/database"
	"sample-api-go/internal/repositories"
	usecase "sample-api-go/internal/use-case"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	dbConnection, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	// camada repo
	SampleRepository := repositories.NewSampleRepository(dbConnection)
	// camada use-case
	SampletUseCase := usecase.NewCreateSampleUseCase(SampleRepository)
	// camada controller
	// endpoinst

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
