package dto

type Merchant struct {
	Id        string      `json:"id"`
	CreatedAt string      `json:"createdAt"`
	Providers []Providers `json:"providers"`
}

type Providers struct {
	Id             string         `json:"id"`
	Type           string         `json:"type"`
	ProviderConfig ProviderConfig `json:"providerConfig"`
}

type ProviderConfig struct {
	PaymentTypes []string `json:"paymentTypes"`
}

type ProvidersOutput struct {
	Id           string   `json:"id"`
	Type         string   `json:"type"`
	PaymentTypes []string `json:"paymentTypes"`
}

type GetMerchantOutput struct {
	Id             string      `json:"id"`
	CreatedAt      string      `json:"createdAt"`
	Providers      []Providers `json:"providers"`
	PaymentMethods []string    `json:"paymentMethods"`
}
