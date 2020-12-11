package bittrex

type Address struct {
	Currency string `json:"Currency"`
	Address  string `json:"Address"`
}

type AddressParams struct {
	CurrencySymbol string `json:"currencySymbol"`
}

type AddressV3 struct {
	Status           string `json:"status"`
	CurrencySymbol   string `json:"currencySymbol"`
	CryptoAddress    string `json:"cryptoAddress"`
	CryptoAddressTag string `json:"cryptoAddressTag"`
}