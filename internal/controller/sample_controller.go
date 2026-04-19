package controllers

import (
	"fmt"
	"net/http"
	apperrors "sample-api-go/internal/errors"
	"sample-api-go/internal/models"
	usecase "sample-api-go/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SampleController struct {
	sampleUseCase usecase.SampleUseCase
}

func NewSampleController(usecase usecase.SampleUseCase) *SampleController {
	return &SampleController{
		sampleUseCase: usecase,
	}
}

func (s *SampleController) GetSamples(ctx *gin.Context) {
	samples, err := s.sampleUseCase.GetSamples()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, samples)
}

func (s *SampleController) GetSampleByID(ctx *gin.Context) {
	id := ctx.Param("id_sample")
	if id == "" {
		response := models.ResponseModel{Message: "Id não pode ser null"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	sampleId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apperrors.BadRequest("id not valid"))
		return
	}

	sample, err := s.sampleUseCase.GetSampleByID(sampleId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if sample == nil {
		ctx.JSON(http.StatusNotFound, apperrors.NotFound(fmt.Sprintf("sample with id %s not found", id)))
		return
	}

	ctx.JSON(http.StatusOK, sample)
}

func (s *SampleController) CreateSample(ctx *gin.Context) {
	var sample models.SampleModel
	err := ctx.BindJSON(&sample)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedSample, err := s.sampleUseCase.CreateSample(sample)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, models.ResponseModel{Message: "Created id " + strconv.Itoa(insertedSample.ID)})
}

func (s *SampleController) SoftDeleteSampleByID(ctx *gin.Context) {
	id := ctx.Param("id_sample")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, apperrors.BadRequest("invalid id, must be not null"))
		return
	}

	sampleId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apperrors.BadRequest("invalid id, must not be char"))
		return
	}

	err = s.sampleUseCase.SoftDeleteSampleByID(sampleId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseModel{Message: "disabled: " + id})
}

func (s *SampleController) HardDeleteSampleByID(ctx *gin.Context) {
	id := ctx.Param("id_sample")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, apperrors.BadRequest("invalid id, must be not null"))
		return
	}

	removeId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apperrors.BadRequest("invalid id, must not be char"))
		return
	}

	err = s.sampleUseCase.HardDeleteSampleByID(removeId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseModel{Message: "removed " + id})
}

func (s *SampleController) UpdateSample(ctx *gin.Context) {
	id := ctx.Param("id_sample")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, apperrors.BadRequest("invalid id, must be not null"))
		return
	}

	sampleId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, apperrors.BadRequest("invalid id, must not be char"))
		return
	}

	var input map[string]interface{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, apperrors.BadRequest("invalid body"))
		return
	}

	updated, err := s.sampleUseCase.UpdateSample(sampleId, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updated)
}
