package bridge

type CryptoApp interface {
	GetCryptoCurrencies(currency string) (interface{}, error)
}
