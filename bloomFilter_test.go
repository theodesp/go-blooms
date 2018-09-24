package go_blooms

import (
	"fmt"
	"hash"
	"hash/fnv"
	"math/rand"
	"sort"
	"testing"

	"github.com/spaolacci/murmur3"
)

// 1 Million
var memberSize uint = 1000000

// Test items if items may exist into set
func TestExistance(t *testing.T) {
	bf := New(memberSize, DefaultHashFunctions)

	for i := 0; i < 1000; i++ {
		item := randomBytes(rand.Intn(54) + 10)

		bf.Add(item)

		if bf.Test(item) != true {
			t.Errorf("'%q' not found", item)
		}

		// Now lets create some items that don't exist
		item2 := append(item, randomBytes(rand.Intn(54)+10)...)

		// Test that item does NOT exist
		if bf.Test(item2) == true {
			t.Errorf("'%q' should not be found", item2)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	bf := New(memberSize, DefaultHashFunctions)
	for i := 0; i < b.N; i++ {
		// Logic to benchmark
		bf.Add([]byte(string(i)))
	}
}

func BenchmarkTest(b *testing.B) {
	bf := New(memberSize, DefaultHashFunctions)
	for i := 0; i < b.N; i++ {
		// Logic to benchmark
		bf.Add([]byte(string(i)))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Logic to benchmark
		bf.Test([]byte(string(i)))
	}
}

func BenchmarkSearch(b *testing.B) {
	bf := New(memberSize, DefaultHashFunctions)
	for i := 0; i < b.N; i++ {
		// Logic to benchmark
		bf.Add([]byte(string(i) + " foo bar baz"))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Logic to benchmark
		bf.Test([]byte(string(i) + " foo baz"))
	}
}

// Compare with a binary sort to make sure we're in the same ballpark
func BenchmarkBinarySearch(b *testing.B) {
	randomString := func() string {
		return fmt.Sprintf("%q", randomBytes(20))
	}

	var strings []string
	for i := 0; i < 100000; i++ {
		item := randomString()
		strings = append(strings, item)
	}

	// Sort by byte order
	sort.Strings(strings)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		item := randomString()
		use(sort.SearchStrings(strings, item))
	}
}

func BenchmarkHashFunctions(b *testing.B) {
	item := []byte("hello world, how are you doing?")

	hashFns := map[string]hash.Hash64{
		"murmur3": murmur3.New64(),
		"fnv64":   fnv.New64(),
		"fnv64a":  fnv.New64a(),
	}

	for name, hashFunc := range hashFns {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				hashFunc.Write(item)
				use(hashFunc.Sum64())
				hashFunc.Reset()
			}
		})
	}
}

func use(interface{}) {}

func randomBytes(size int) []byte {
	b := make([]byte, size)
	rand.Read(b)
	return b
}
