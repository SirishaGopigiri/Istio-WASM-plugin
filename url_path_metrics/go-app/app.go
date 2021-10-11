package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

type testHeader struct {
	Username   string    `header:"Username"`
}

func RunServer() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world!!",
		})
	})
	r.GET("/login", func(c *gin.Context) {
		h := testHeader{}
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(200, err)
		}
		if h.Username == "" {
			c.JSON(500, "Expected Username!!")
			return
		}
		fmt.Printf("%#v\n", h)

		c.JSON(200, gin.H{
			"message": "hello "+h.Username+"!! Login Successful!!",
		})
	})
	return r
}

func main() {
	RunServer().Run(":5000")
}

