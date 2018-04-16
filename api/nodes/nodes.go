package nodes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cbarraford/marvel-coin/node"
)

func List() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, node.FilterList())
	}
}
