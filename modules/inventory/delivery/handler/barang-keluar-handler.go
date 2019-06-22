package handler

import (
	"log"
	"svc-inventory-go/modules/inventory/service"

	"github.com/gin-gonic/gin"
)

func (dependencyInjection *DatabaseInjection) GetAllBarangKeluar(ctx *gin.Context) {
	allBarangKeluarService, errorHandlerService := service.FetchAllBarangKeluar(dependencyInjection.DB)

	if errorHandlerService != nil {
		log.Fatalf("Error exception %v", errorHandlerService)
	}

	if allBarangKeluarService != nil {
		ctx.JSON(200, gin.H{
			"total": len(allBarangKeluarService),
			"count": len(allBarangKeluarService),
			"data":  allBarangKeluarService,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Data barang keluar kosong",
		})
	}

}

func (dependencyInjection *DatabaseInjection) GetDetailBarangKeluar(ctx *gin.Context) {
	IDBarangKeluar := ctx.Param("IDBarangKeluar")

	detailBarangKeluarService, errorHandlerService := service.FetchDetailBarangKeluar(IDBarangKeluar, dependencyInjection.DB)

	if errorHandlerService != nil {
		log.Fatalf("Error exception %v", errorHandlerService)
	}

	if detailBarangKeluarService != nil {
		ctx.JSON(200, gin.H{
			"data": detailBarangKeluarService,
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "Data barang keluar kosong",
		})
	}

}
