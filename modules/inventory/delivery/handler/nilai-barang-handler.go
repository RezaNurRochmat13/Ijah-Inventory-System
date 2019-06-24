package handler

import (
	"log"
	"svc-inventory-go/modules/inventory/service"

	"github.com/gin-gonic/gin"
)

func (dependencyInjection *DatabaseInjection) GetTotalLaba(ctx *gin.Context) {
	totalLabaService, errorHandlerService := service.FetchTotalLaba(dependencyInjection.DB)

	if errorHandlerService != nil {
		log.Printf("Error exception %v", errorHandlerService)
	}

	if totalLabaService != nil {
		ctx.JSON(200, gin.H{
			"data": totalLabaService,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Data kosong",
		})
	}
}
