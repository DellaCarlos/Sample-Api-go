package controllers

import (
	"net/http"
	"sample-api-go/internal/models"
	"sample-api-go/internal/usecase"

	"github.com/gin-gonic/gin"
)

type SectorController struct {
	sectorUseCase usecase.SectorUseCase
}

func NewSectorController(usecase usecase.SectorUseCase) *SectorController {
	return &SectorController{
		sectorUseCase: usecase,
	}
}

func (sc *SectorController) GetSector(ctx *gin.Context) {
	sactors, err := sc.sectorUseCase.GetSectors()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, sactors)
}

func (sc *SectorController) CreateSector(ctx *gin.Context) {
	var sector models.SectorModel
	err := ctx.BindJSON(&sector)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedSector, err := sc.sectorUseCase.CreateSector(sector)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, models.ResponseModel{Message: "Created: " + insertedSector.Sector})
}
