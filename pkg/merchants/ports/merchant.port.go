package http_port

import (
	"context"
	"encoding/json"
	"merchants/pkg/merchants/dto"
	"net/http"
	"os"
)

type ProviderWS struct {
	Name      string          `json:"name"`
	Id        string          `json:"id"`
	CreatedAt string          `json:"createdAt"`
	Providers []dto.Providers `json:"providers"`
}

type MerchantApi struct {
}

func NewMerchantApi() *MerchantApi {
	return &MerchantApi{}
}

type ParamsMerchant struct {
	MerchantId string
}

func (m *MerchantApi) Request(ctx context.Context, params ParamsMerchant) (*ProviderWS, error) {
	resp, err := http.Get(os.Getenv("MERCHANTS_URL") + params.MerchantId)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var provider ProviderWS
	err = json.NewDecoder(resp.Body).Decode(&provider)
	if err != nil {
		return nil, err
	}

	return &provider, nil
}
