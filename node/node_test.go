package node

import (
	"testing"
	"time"

	check "gopkg.in/check.v1"
)

func TestPackage(t *testing.T) { check.TestingT(t) }

type NodeSuite struct{}

var _ = check.Suite(&NodeSuite{})

func (s *NodeSuite) TestAdd(c *check.C) {
	Add(Node{DNS: "test.marvel.com", IP: "1.2.3.4", LastSeen: time.Now()})
	c.Check(nodes, check.HasLen, 1)

	// Do not allow the addition of duplicates
	Add(Node{DNS: "test.marvel.com", IP: "1.2.3.4", LastSeen: time.Now()})
	c.Check(nodes, check.HasLen, 1)

	// Do not allow adding an old node
	Add(Node{DNS: "test2.marvel.com", IP: "1.2.3.5", LastSeen: time.Now().Add(-100 * time.Hour)})
	c.Check(nodes, check.HasLen, 1)

	// Do not allow adding future last seen
	Add(Node{DNS: "test2.marvel.com", IP: "1.2.3.5", LastSeen: time.Now().Add(100 * time.Hour)})
	c.Check(nodes, check.HasLen, 1)
}

func (s *NodeSuite) TestFilter(c *check.C) {
	now := time.Now()
	ns := []Node{
		Node{DNS: "test.marvel.com", IP: "1.2.3.4", LastSeen: now},
		Node{DNS: "test.marvel.com", IP: "1.2.3.4", LastSeen: now},
		Node{DNS: "test2.marvel.com", IP: "1.2.3.5", LastSeen: now.Add(-100 * time.Hour)},
	}

	ns = Filter(ns)
	c.Assert(ns, check.HasLen, 1)
	c.Check(ns[0].IP, check.Equals, "1.2.3.4")
	c.Check(ns[0].DNS, check.Equals, "test.marvel.com")
	c.Check(ns[0].LastSeen.Unix(), check.Equals, now.Unix())
}

func (s *NodeSuite) TestList(c *check.C) {
	now := time.Now()
	nodes = []Node{
		Node{DNS: "test.marvel.com", IP: "1.2.3.4", LastSeen: now},
		Node{DNS: "test.marvel.com", IP: "1.2.3.4", LastSeen: now},
		Node{DNS: "test2.marvel.com", IP: "1.2.3.5", LastSeen: now.Add(-100 * time.Hour)},
	}

	c.Assert(List(), check.HasLen, 3)
	c.Assert(FilterList(), check.HasLen, 1)
}
