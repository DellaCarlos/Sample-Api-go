package main

import "github.com/gin-gonic/gin"

func SamplesRoutes(router *gin.Engine) {
	samplesRoutes := router.Group("/samples")

	// Initialize the postgresql repository
	// psqlCategoryRepository :=

	samplesRoutes.GET("/", func(ctx *gin.Context) {
		// TODO: Implement the logic to get all samples
	})

	samplesRoutes.POST("/", func(ctx *gin.Context) {
		// TODO: Implement the logic to create a new sample
	})
}
