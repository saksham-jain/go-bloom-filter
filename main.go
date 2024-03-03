package main

import (
	"fmt"
	bf "sj_bloom_filter/bloom_filter"
	hs "sj_bloom_filter/hash_strategy"
)

func main() {
	var isPresent bool
	size := uint32(4)
	hash_strategy := hs.New(size)
	bloomFilter := bf.New(size, hash_strategy)
	bloomFilter.Set("Banana")
	fmt.Printf("bloomFilter: %+v\n", bloomFilter)
	isPresent, _ = bloomFilter.IsPresent("Banana")

	fmt.Printf("Banana is present? : %v\n", isPresent)

	bloomFilter.Set("Apple")
	fmt.Printf("bloomFilter: %+v\n", bloomFilter)
	isPresent, _ = bloomFilter.IsPresent("Apple")
	fmt.Printf("Apple is present? : %v\n", isPresent)

	isPresent, _ = bloomFilter.IsPresent("grrr") // false positive
	fmt.Printf("grrr is present? : %v\n", isPresent)
}

// code response example--
// MurmurHash3 hash value: 1
// bloomFilter: &{hashStrategy:0x1400000c030 bitArr:[0 1 0 0]}
// MurmurHash3 hash value: 1
// Banana is present? : true
// MurmurHash3 hash value: 2
// bloomFilter: &{hashStrategy:0x1400000c030 bitArr:[0 1 1 0]}
// MurmurHash3 hash value: 2
// Apple is present? : true
// MurmurHash3 hash value: 2
// grrr is present? : true
