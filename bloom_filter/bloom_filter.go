package bloom_filter

import (
	"sj_bloom_filter/hash_strategy"
)

type bloomFilter struct {
	hashStrategy hash_strategy.HashStrategyInterface
	bitArr       []uint32
}

func New(size uint32, hashStrategy hash_strategy.HashStrategyInterface) FilterInterface {
	bitArr := make([]uint32, size)
	return &bloomFilter{
		hashStrategy: hashStrategy,
		bitArr:       bitArr,
	}
}

func (bf *bloomFilter) IsPresent(element string) (bool, error) {
	key, err := bf.GetHashKey(element)
	if err != nil {
		return false, err
	}

	if bf.bitArr[*key] == 1 {
		return true, nil
	} else {
		return false, nil
	}

}

func (bf *bloomFilter) Set(element string) error {
	key, err := bf.GetHashKey(element)
	if err != nil {
		return err
	}

	bf.bitArr[*key] = 1

	return nil
}

func (bf *bloomFilter) GetHashKey(element string) (*uint32, error) {
	key, err := bf.hashStrategy.GetKey(element)
	if err != nil {
		return nil, err
	}
	return &key, nil
}
