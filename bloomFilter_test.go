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
