package go_blooms

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestTrue(c *C) {
	c.Assert(true, Equals, true)
}

// Basic init test
func (s *MySuite) TestNew(c *C)  {
	bf := New(1024, 3)

	c.Assert(bf.k, Equals, uint(3))
	c.Assert(bf.n, Equals, uint(0))
	c.Assert(bf.m, Equals, uint(1024))
}

// Test add to set
func (s *MySuite) TestAdd(c *C)  {
	bf := New(1024, 3)

	bf.Add([]byte("hello"))

	c.Assert(bf.n, Equals, uint(1))
}

// Test items if do not exist into set
func (s *MySuite) TestIfNotExist(c *C)  {
	bf := New(1024, 3)

	bf.Add([]byte("hello"))
	bf.Add([]byte("world"))
	bf.Add([]byte("hi"))

	c.Assert(bf.Test([]byte("H1")), Equals, false)
	c.Assert(bf.Test([]byte("World")), Equals, false)
	c.Assert(bf.Test([]byte("hell0")), Equals, false)
}

// Test items if items may exist into set
func (s *MySuite) TestIfMayExist(c *C)  {
	bf := New(6, 3)

	bf.Add([]byte("hello"))
	bf.Add([]byte("world"))
	bf.Add([]byte("sir"))
	bf.Add([]byte("madam"))
	bf.Add([]byte("io"))

	c.Assert(bf.Test([]byte("hello")), Equals, true)
	c.Assert(bf.Test([]byte("world")), Equals, true)
	// False negative
	c.Assert(bf.Test([]byte("hi")), Equals, false)
}
