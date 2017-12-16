go-blooms
---

From Wiki
>Bloom Filter: A space-efficient probabilistic data structure that is used to test whether an element is a member of 
a set. False positive matches are possible, but false negatives are not; i.e. a query returns either "possibly in set" 
or "definitely not in set". Elements can be added to the set, but not removed.

This bloom filter implementation is backed by bool slice for simplicity.

And the hashing functions used are fnv and murmur both 64 bit versions.

## Installation

go get -u github.com/theodesp/go-blooms

## Usage

```go
package example
import "github.com/theodesp/go-blooms"

const (
  size = 64 * 1024
  hashFunctionsSize = 3
)

bf := bloomfilter.New(size, hashFunctionsSize)

value := "hello"

bf.Add([]byte(value)) // we accept only a byte slice
if bf.Test(value) { // probably true, could be false
  // whatever
}

anotherValue := "world"

if bf.Test(anotherValue) { // Bloom filter guarantees that anotherValue is not in the set
  panic("This should never happen")
}

```

## Complexity

**Time**

If we are using a bloom filter with  bits and  hash function, 
insertion and search will both take  time. 
In both cases, we just need to run the input through all of 
the hash functions. Then we just check the output bits.

|  Operation | Complexity  |
|---|---|
|  insertion |  O(k)  |
|  search |  O(k)  |

**Space**
The space of the actual data structure (what holds the data).

|  Complexity |
|---|
|  O(m) |

Where `m` is the size of the slice.


## License

Copyright © 2017 Theo Despoudis
MIT license