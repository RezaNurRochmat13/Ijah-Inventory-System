package product

import (
	"svc-inventory-go/modules/inventory/delivery/handler"
	"svc-inventory-go/resources"

	"github.com/gin-gonic/gin"
)

func InventoryRoutes(routes *gin.Engine) {

	databaseConnInit := resources.SetupDatabaseConfiguration()

	dependencyInjection := &handler.DatabaseInjection{
		DB: databaseConnInit,
	}

	inventoryRouteConfiguration := routes.Group("/public/api/v1")
	{
		inventoryRouteConfiguration.GET("product-income", dependencyInjection.GetAllBarangMasuk)
	}
}
