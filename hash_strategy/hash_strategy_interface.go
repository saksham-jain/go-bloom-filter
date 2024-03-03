package hash_strategy

type HashStrategyInterface interface {
	GetKey(value interface{}) (uint32, error)
}
