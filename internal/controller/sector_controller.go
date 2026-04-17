package controllers

import (
	"net/http"
	apperrors "sample-api-go/internal/errors"
	"sample-api-go/internal/models"
	"sample-api-go/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	ctx.JSON(http.StatusCreated, models.ResponseModel{Message: "Sector created: " + insertedSector.Sector})
}

func (sc *SectorController) DeleteSector(ctx *gin.Context) {
	id := ctx.Param("id_sector")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, apperrors.BadRequest("invalid id, must be not null"))
		return
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apperrors.BadRequest("invalid id, must be a valid UUID"))
		return
	}

	err = sc.sectorUseCase.DeleteSector(parsedID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseModel{Message: "removed " + id})
}
