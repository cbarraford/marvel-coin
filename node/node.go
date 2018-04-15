package node

import (
	"time"
)

// TODO add function to check if nodes are still available

const expireIn = 3600

var nodes []Node

type Node struct {
	DNS      string    `json:"dns"`
	IP       string    `json:"ip"`
	LastSeen time.Time `json:"last_seen"`
}

// Add a new node to the node list. If node is expired or already exists, it
// will not be added.
func Add(node Node) {
	if !isFuture(node) && !isExpired(node) && !doesExist(nodes, node) {
		nodes = append(nodes, node)
	}
}

// Filter the given list of nodes, removing expired and duplicates
func Filter(ns []Node) []Node {
	newList := []Node{}

	for _, n := range ns {
		if !isExpired(n) && !doesExist(newList, n) {
			newList = append(newList, n)
		}
	}

	return newList
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

// Return list of nodes
func List() []Node {
	return nodes
}

// Return list of filter nodes
func FilterList() []Node {
	return Filter(nodes)
}
