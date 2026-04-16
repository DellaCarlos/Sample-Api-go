package routes

import (
	controllers "sample-api-go/internal/controller"

	"github.com/gin-gonic/gin"
)

func Register(server *gin.Engine, sampleController *controllers.SampleController, sectorController *controllers.SectorController) {
	v1 := server.Group("/api/v1")
	{
		samples := v1.Group("/samples")
		{
			samples.GET("", sampleController.GetSamples)
			samples.GET("/:id_sample", sampleController.GetSampleByID)
			samples.POST("", sampleController.CreateSample)
			samples.DELETE("/d/:id_sample", sampleController.SoftDeleteSampleByID)
			samples.DELETE("/hd/:id_sample", sampleController.HardDeleteSampleByID)
		}

		sectors := v1.Group("/sectors")
		{
			sectors.GET("", sectorController.GetSector)
			sectors.POST("", sectorController.CreateSector)
		}
	}
}
