package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "gopkg.in/check.v1"
)

func TestPackage(t *testing.T) { TestingT(t) }

type ApiSuite struct{}

var _ = Suite(&ApiSuite{})

func (s *ApiSuite) TestApiService(c *C) {
	r := GetAPIService()

	// check ping apiendpoint
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	c.Assert(w.Code, Equals, 200)
}
