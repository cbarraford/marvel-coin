package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cbarraford/marvel-coin/api/nodes"
	"github.com/cbarraford/marvel-coin/node"
)

func GetAPIService(nodeSet *node.NodeSet) *gin.Engine {
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Availability API Health Check
	r.GET("/ping", ping())

	r.GET("/nodes", nodes.List(nodeSet))

	return r
}

// health-check to test service is up
func ping() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
