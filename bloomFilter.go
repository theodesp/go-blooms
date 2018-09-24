package go_blooms

import (
	"hash"
	"hash/fnv"

	"github.com/spaolacci/murmur3"
)

// Minimal interface that the Bloom filter must implement
type Interface interface {
	Add(item []byte)       // Adds the item into the Set
	Test(item []byte) bool // Performs probabilist test if the item exists in the set or not.
}

// BloomFilter probabilistic data structure definition
type BloomFilter struct {
	bitset  []bool        // The bloom-filter bitset
	n       uint          // Number of elements in the filter
	m       uint          // Size of the bloom filter
	hashfns []hash.Hash64 // The hash functions
}

// DefaultHashFunctions for BloomFilter
var DefaultHashFunctions = []hash.Hash64{murmur3.New64(), fnv.New64(), fnv.New64a()}

// Returns a new BloomFilter object,
func New(size uint, hashes []hash.Hash64) *BloomFilter {
	return &BloomFilter{
		bitset:  make([]bool, size),
		m:       size,
		n:       uint(0),
		hashfns: hashes,
	}
}

// Adds the item into the bloom filter set by hashing in over the hash functions
func (bf *BloomFilter) Add(item []byte) {
	for _, v := range bf.hashValues(item) {
		position := uint(v) % bf.m
		bf.bitset[position] = true
	}
	bf.n += 1
}

// Test if the item into the bloom filter is set by hashing in over the hash functions
func (bf *BloomFilter) Test(item []byte) bool {
	for _, v := range bf.hashValues(item) {
		position := uint(v) % bf.m
		if !bf.bitset[uint(position)] {
			return false
		}
	}
	return true
}

// Calculates all the hash values by hashing in over the hash functions
func (bf *BloomFilter) hashValues(item []byte) []uint64 {
	var result []uint64

	for _, hashFunc := range bf.hashfns {
		hashFunc.Write(item)
		result = append(result, hashFunc.Sum64())
		hashFunc.Reset()
	}

	return result
}
