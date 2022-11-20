package bridge

type App interface {
	GetCryptoCurrencies(currency string) (interface{}, error)
	FindHash(hashStr string) (interface{}, error)
}
