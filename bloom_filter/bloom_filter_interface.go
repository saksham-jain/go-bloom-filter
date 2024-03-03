package bloom_filter

type FilterInterface interface {
	IsPresent(element string) (bool, error)
	Set(element string) error
}
