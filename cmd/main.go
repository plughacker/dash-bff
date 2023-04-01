package main

import (
	"merchants/pkg/merchants"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controller := merchants.NewMerchantController()
	r.GET("/v1/merchant", func(c *gin.Context) {
		id := c.Param("id")
		response := controller.GetMerchantWithProviderAndPaymentMethod(c, id)
		c.JSON(200, response)
	})
	r.Run()
}
