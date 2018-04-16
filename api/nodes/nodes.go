package nodes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cbarraford/marvel-coin/node"
)

func List(nodeSet *node.NodeSet) func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, nodeSet.FilterList())
	}
}
