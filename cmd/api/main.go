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

	SampleRepository := repositories.NewSampleRepository(dbConnection)  // camada repo
	SampletUseCase := usecase.NewCreateSampleUseCase(SampleRepository)  // camada use-case
	SampleController := controllers.NewSampleController(SampletUseCase) // camada controller

	// endpoinst
	server.GET("/samples", SampleController.GetSamples)
	server.GET("/samples/:id_sample", SampleController.GetSampleByID)
	server.POST("/sample", SampleController.CreateSample)
	server.DELETE("/sample/:id_sample", SampleController.SoftDeleteSampleByID)

	// Start server on port 8080 (default)
	server.Run()
}
