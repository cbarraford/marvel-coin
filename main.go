package main

import (
	"github.com/cbarraford/marvel-coin/api"
	"github.com/cbarraford/marvel-coin/node"
)

func main() {
	nodeSet := node.NewNodeSet()

	serve := api.GetAPIService(nodeSet)
	serve.Run()
}
