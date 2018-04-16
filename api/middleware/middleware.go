package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func TestSuite() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // execute all the handlers

		errorToPrint := c.Errors.Last()
		if errorToPrint != nil {
			log.Printf("API ERROR: %+v", errorToPrint.Error())
		}
	}
}
