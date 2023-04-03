package merchants

import (
	"github.com/gin-gonic/gin"
)

type MerchantController struct {
}

func NewMerchantController() *MerchantController {
	return &MerchantController{}
}

func (m *MerchantController) GetMerchantWithProviderAndPaymentMethod(c *gin.Context, id string) any {
	merchantService := NewGetMerchant()
	merchants := merchantService.GetMerchant(c, id)

	return merchants
}
