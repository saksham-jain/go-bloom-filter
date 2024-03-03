package hash_strategy

import (
	"fmt"
	"hash"
	"strconv"

	"github.com/spaolacci/murmur3"
)

type murMurHashStrategy struct {
	size uint32
	hash hash.Hash32
}

func New(size uint32) HashStrategyInterface {
	return &murMurHashStrategy{
		size: size,
		hash: murmur3.New32(),
	}
}

func (m *murMurHashStrategy) GetKey(value interface{}) (uint32, error) {
	m.hash.Reset()

	switch v := value.(type) {
	case int:
		str := strconv.Itoa(v)
		byteSlice := []byte(str)
		m.hash.Write(byteSlice)
	case string:
		byteSlice := []byte(v)
		m.hash.Write(byteSlice)
	}

	hashValue := m.hash.Sum32()
	finalValue := hashValue % m.size
	fmt.Printf("MurmurHash3 hash value: %d\n", finalValue)
	return finalValue, nil
}
