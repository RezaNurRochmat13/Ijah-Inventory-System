package handler

import "github.com/gin-gonic/gin"

func GetAllStock(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "All stock product",
	})
}
