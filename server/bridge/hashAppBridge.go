package bridge

type HashApp interface {
	FindHash(hashStr string) (interface{}, error)
}
