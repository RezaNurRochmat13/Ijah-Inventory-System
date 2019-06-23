package handler

import (
	"log"
	"svc-inventory-go/modules/inventory/service"

	"github.com/gin-gonic/gin"
)

func (dependencyInjection *DatabaseInjection) GetAllStokBarang(ctx *gin.Context) {
	resultStokBarangService, errorHandlerService := service.FetchAllStokBarang(dependencyInjection.DB)

	if errorHandlerService != nil {
		log.Printf("Error exception %v", errorHandlerService)
	}

	if resultStokBarangService != nil {
		ctx.JSON(200, gin.H{
			"total": len(resultStokBarangService),
			"data":  resultStokBarangService,
			"count": len(resultStokBarangService),
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Data stok barang kosong",
		})
	}
}
