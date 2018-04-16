package node

import (
	"fmt"
	"net/http"
	"time"
)

const expireIn = 3600

type NodeSet struct {
	Nodes []Node
	http  http.Client
}

type Node struct {
	DNS      string    `json:"dns"`
	IP       string    `json:"ip"`
	LastSeen time.Time `json:"last_seen"`
}

func NewNodeSet() *NodeSet {
	return &NodeSet{
		Nodes: nil,
		http:  http.Client{},
	}
}

// Add a new node to the node list. If node is expired or already exists, it
// will not be added.
func (s *NodeSet) Add(node Node) {
	if !isFuture(node) && !isExpired(node) && !doesExist(s.Nodes, node) {
		s.Nodes = append(s.Nodes, node)
	}
}

// Filter the given list of nodes, removing expired and duplicates
func (s *NodeSet) Filter(ns []Node) []Node {
	newList := []Node{}

	for _, n := range ns {
		if !isExpired(n) && !doesExist(newList, n) {
			newList = append(newList, n)
		}
	}

	return newList
}

// Return list of nodes
func (s *NodeSet) List() []Node {
	return s.Nodes
}

// Return list of filter nodes
func (s *NodeSet) FilterList() []Node {
	return s.Filter(s.Nodes)
}

func (s *NodeSet) Check(node Node) bool {
	resp, err := s.http.Get(fmt.Sprintf("http://%s/ping", node.IP))
	if err != nil {
		return false
	}
	return resp.StatusCode == 200
}

// Check if time of last seen is expired
func isExpired(node Node) bool {
	diff := time.Now().Unix() - node.LastSeen.Unix()
	return diff > expireIn
}

// Check if timestamp is the future
func isFuture(node Node) bool {
	diff := time.Now().Unix() - node.LastSeen.Unix()
	return diff < 0
}

func doesExist(ns []Node, node Node) bool {
	for _, n := range ns {
		if n.IP == node.IP && n.DNS == node.DNS {
			return true
		}
	}
	return false
}
