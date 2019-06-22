package product

import (
	"svc-inventory-go/modules/inventory/delivery/handler"

	"github.com/gin-gonic/gin"
)

func InventoryRoutes(routes *gin.Engine) {

	inventoryRouteConfiguration := routes.Group("/public/api/v1")
	{
		inventoryRouteConfiguration.GET("stock-product", handler.GetAllStock)
	}
}
