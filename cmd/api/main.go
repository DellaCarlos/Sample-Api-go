package main

import (
	controllers "sample-api-go/internal/controller"
	"sample-api-go/internal/database"
	"sample-api-go/internal/repositories"
	usecase "sample-api-go/internal/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// middleware
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"}, // porta vite
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	dbConnection, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	SampleRepository := repositories.NewSampleRepository(dbConnection)
	SampletUseCase := usecase.NewCreateSampleUseCase(SampleRepository)
	SampleController := controllers.NewSampleController(SampletUseCase)

	server.GET("/samples", SampleController.GetSamples)
	server.GET("/samples/:id_sample", SampleController.GetSampleByID)
	server.POST("/sample", SampleController.CreateSample)
	server.DELETE("/sampledelete/:id_sample", SampleController.SoftDeleteSampleByID)

	server.Run()
}
