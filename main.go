package main

import (
	"github.com/cbarraford/marvel-coin/api"
)

func main() {
	serve := api.GetAPIService()
	serve.Run()
}
