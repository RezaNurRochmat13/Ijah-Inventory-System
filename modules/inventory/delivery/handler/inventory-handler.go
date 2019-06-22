package handler

import (
	"log"
	"svc-inventory-go/modules/inventory/service"

	"github.com/gin-gonic/gin"
)

func (dependencyInjection *DatabaseInjection) GetAllBarangMasuk(ctx *gin.Context) {

	barangMasukResultFromService, errorHandlerService := service.FetchAllBarangMasukService(dependencyInjection.DB)

	if errorHandlerService != nil {
		log.Fatalf("Error exception %v", errorHandlerService)
	}

	if barangMasukResultFromService != nil {
		ctx.JSON(200, gin.H{
			"total": len(barangMasukResultFromService),
			"data":  barangMasukResultFromService,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Data barang masuk kosong",
		})
	}

}
