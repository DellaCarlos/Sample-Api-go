package main

import (
	controllers "sample-api-go/internal/controller"
	"sample-api-go/internal/database"
	"sample-api-go/internal/repositories"
	usecase "sample-api-go/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	// camada repo
	SampleRepository := repositories.NewSampleRepository(dbConnection)
	// camada use-case
	SampletUseCase := usecase.NewCreateSampleUseCase(SampleRepository)
	// camada controller
	SampleController := controllers.NewSampleController(SampletUseCase)

	// endpoinst
	server.GET("/samples", SampleController.GetSamples)
	server.GET("/samples/:sampleId", SampleController.GetSampleByID)
	server.POST("/sample", SampleController.CreateSample)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	server.Run()
}
