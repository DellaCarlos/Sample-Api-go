package main

import (
	"log"
	controllers "sample-api-go/internal/controller"
	"sample-api-go/internal/database"
	"sample-api-go/internal/repositories"
	"sample-api-go/internal/routes"
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
		log.Fatalf("erro ao conectar no banco: %v", err)
	}

	sampleRepository := repositories.NewSampleRepository(dbConnection)
	sampleUseCase := usecase.NewCreateSampleUseCase(sampleRepository)
	sampleController := controllers.NewSampleController(sampleUseCase)

	// Sectors
	sectorRepository := repositories.NewSectorRepository(dbConnection)
	sectorUseCase := usecase.NewSectorUseCase(sectorRepository)
	sectorController := controllers.NewSectorController(sectorUseCase)

	routes.Register(server, sampleController, sectorController)

	if err := server.Run(":8000"); err != nil {
		log.Fatalf("erro ao iniciar servidor: %v", err)
	}
}
