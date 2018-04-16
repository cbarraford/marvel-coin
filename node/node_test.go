package node

import (
	"net/http"
	"testing"
	"time"

	check "gopkg.in/check.v1"
)

func TestPackage(t *testing.T) { check.TestingT(t) }

type NodeSuite struct{}

var _ = check.Suite(&NodeSuite{})

type ClientMock struct {
}

func (c *ClientMock) Get(url string) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
	}, nil
}

type BadClientMock struct {
}

func (c *BadClientMock) Get(url string) (*http.Response, error) {
	return &http.Response{
		StatusCode: 404,
	}, nil
}

func (s *NodeSuite) TestAdd(c *check.C) {
	nodeSet := NewNodeSet()
	nodeSet.Add(Node{DNS: "test.marvel.com", IP: "1.2.3.4", LastSeen: time.Now()})
	c.Assert(nodeSet.Nodes, check.HasLen, 1)

	// Do not allow the addition of duplicates
	nodeSet.Add(Node{DNS: "test.marvel.com", IP: "1.2.3.4", LastSeen: time.Now()})
	c.Check(nodeSet.Nodes, check.HasLen, 1)

	// Do not allow adding an old node
	nodeSet.Add(Node{DNS: "test2.marvel.com", IP: "1.2.3.5", LastSeen: time.Now().Add(-100 * time.Hour)})
	c.Check(nodeSet.Nodes, check.HasLen, 1)

	// Do not allow adding future last seen
	nodeSet.Add(Node{DNS: "test2.marvel.com", IP: "1.2.3.5", LastSeen: time.Now().Add(100 * time.Hour)})
	c.Check(nodeSet.Nodes, check.HasLen, 1)
}

func (s *NodeSuite) TestRemove(c *check.C) {
	nodeSet := NewNodeSet()
	n := Node{DNS: "test.marvel.com", IP: "1.2.3.4", LastSeen: time.Now()}
	nodeSet.Add(n)
	c.Assert(nodeSet.Nodes, check.HasLen, 1)

	nodeSet.Remove(n)
	c.Assert(nodeSet.Nodes, check.HasLen, 0)
}

func (s *NodeSuite) TestFilter(c *check.C) {
	nodeSet := NewNodeSet()
	now := time.Now()
	ns := []Node{
		Node{DNS: "test.marvel.com", IP: "1.2.3.4", LastSeen: now},
		Node{DNS: "test.marvel.com", IP: "1.2.3.4", LastSeen: now},
		Node{DNS: "test2.marvel.com", IP: "1.2.3.5", LastSeen: now.Add(-100 * time.Hour)},
	}

	ns = nodeSet.Filter(ns)
	c.Assert(ns, check.HasLen, 1)
	c.Check(ns[0].IP, check.Equals, "1.2.3.4")
	c.Check(ns[0].DNS, check.Equals, "test.marvel.com")
	c.Check(ns[0].LastSeen.Unix(), check.Equals, now.Unix())
}

func (s *NodeSuite) TestList(c *check.C) {
	nodeSet := NewNodeSet()
	now := time.Now()
	nodeSet.Nodes = []Node{
		Node{DNS: "test.marvel.com", IP: "1.2.3.4", LastSeen: now},
		Node{DNS: "test.marvel.com", IP: "1.2.3.4", LastSeen: now},
		Node{DNS: "test2.marvel.com", IP: "1.2.3.5", LastSeen: now.Add(-100 * time.Hour)},
	}

	c.Assert(nodeSet.List(), check.HasLen, 3)
	c.Assert(nodeSet.FilterList(), check.HasLen, 1)
}

func (s *NodeSuite) TestCheck(c *check.C) {
	nodeSet := NewNodeSet()
	n := Node{DNS: "test.marvel.com", IP: "1.2.3.4", LastSeen: time.Now()}

	nodeSet.http = &ClientMock{}
	c.Assert(nodeSet.Check(n), check.Equals, true)
	nodeSet.http = &BadClientMock{}
	c.Assert(nodeSet.Check(n), check.Equals, false)
}
