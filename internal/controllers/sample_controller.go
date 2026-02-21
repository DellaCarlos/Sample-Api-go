package controllers

import (
	"net/http"
	"sample-api-go/internal/models"
	usecase "sample-api-go/internal/use-case"
	"strconv"

	"github.com/gin-gonic/gin"
)

type sampleController struct {
	sampleUseCase usecase.SampleUseCase
}

func NewSampleController(usecase usecase.SampleUseCase) sampleController {
	return sampleController{
		sampleUseCase: usecase,
	}
}

// função que trata a requisição de obtenção de produtos (get)
func (s *sampleController) GetSamples(ctx *gin.Context) {
	samples, err := s.sampleUseCase.GetSamples()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, samples)
}

// função que trata a requisição de obtenção de produtos (get)
func (s *sampleController) GetSampleByID(ctx *gin.Context) {
	id := ctx.Param("sampleId")
	if id == "" {
		response := models.ResponseModel{
			Message: "Id não pode ser null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	sampleId, err := strconv.Atoi(id)
	if err != nil {
		response := models.ResponseModel{
			Message: "Id precisa ser um número inteiro",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	sample, err := s.sampleUseCase.GetSampleByID(sampleId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if sample == nil {
		response := models.ResponseModel{
			Message: "Produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, sample)
}

// função que trata a requisição de criação de produto (post)
func (s *sampleController) CreateSample(ctx *gin.Context) {
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

	ctx.JSON(http.StatusCreated, insertedSample)
}
