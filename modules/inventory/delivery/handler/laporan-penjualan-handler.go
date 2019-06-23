package handler

import (
	"log"
	"svc-inventory-go/modules/inventory/service"

	"github.com/gin-gonic/gin"
)

func (dependencyInjection *DatabaseInjection) GetAllPenjualan(ctx *gin.Context) {
	resultPenjualanService, errorHandlerService := service.FetchAllPenjualan(dependencyInjection.DB)

	if errorHandlerService != nil {
		log.Printf("Error exception %v", errorHandlerService)
	}

	if resultPenjualanService != nil {
		ctx.JSON(200, gin.H{
			"total": len(resultPenjualanService),
			"data":  resultPenjualanService,
			"count": len(resultPenjualanService),
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Data penjualan kosong",
		})
	}
}
