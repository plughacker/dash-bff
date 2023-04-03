package merchants

import (
	"fmt"
	"merchants/pkg/merchants/dto"
	http_port "merchants/pkg/merchants/ports"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type GetMerchant struct {
}

func NewGetMerchant() *GetMerchant {
	return &GetMerchant{}
}

func (g *GetMerchant) GetMerchant(c *gin.Context, id string) any {
	fmt.Print("entrei! :D")
	response, _ := http_port.NewMerchantApi().Request(c, http_port.ParamsMerchant{
		MerchantId: id,
	})
	fmt.Print("response", response)
	merchant := dto.Merchant{
		Id:        "idMerchant",
		CreatedAt: "2023-03-31 10:05Z",
		Providers: []dto.Providers{{
			Id:   "idProvider",
			Type: "SANDBOX",
			ProviderConfig: dto.ProviderConfig{
				PaymentTypes: []string{"pix", "credit", "boleto"},
			},
		}, {
			Id:   "idProvider2",
			Type: "CIELO",
			ProviderConfig: dto.ProviderConfig{
				PaymentTypes: []string{"pix", "boleto"},
			},
		}},
	}

	providers := g.GetProviders(c, merchant)
	paymentMethods := g.GetPaymentMethods(c, providers)
	return dto.GetMerchantOutput{
		PaymentMethods: paymentMethods,
		Providers:      providers,
		Id:             merchant.Id,
		CreatedAt:      merchant.CreatedAt,
	}
}

func (g *GetMerchant) GetProviders(c *gin.Context, merchant dto.Merchant) []dto.Providers {
	return merchant.Providers
}

func (g *GetMerchant) GetPaymentMethods(c *gin.Context, providers []dto.Providers) []string {
	var paymentMethods []string
	for _, element := range providers {
		for _, value := range element.ProviderConfig.PaymentTypes {
			paymentMethods = append(paymentMethods, value)
		}
	}
	return lo.Uniq(paymentMethods)
}
