package nodes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	check "gopkg.in/check.v1"

	"github.com/cbarraford/marvel-coin/api/middleware"
	"github.com/cbarraford/marvel-coin/node"
)

func TestPackage(t *testing.T) { check.TestingT(t) }

type NodesSuite struct{}

var _ = check.Suite(&NodesSuite{})

func (s *NodesSuite) TestAssign(c *check.C) {
	gin.SetMode(gin.ReleaseMode)

	set := node.NewNodeSet()

	r := gin.New()
	r.Use(middleware.TestSuite())
	r.GET("/nodes", List(set))
	req, _ := http.NewRequest("GET", "/nodes", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	c.Assert(w.Code, check.Equals, 200)
}
