package main

import "github.com/gin-gonic/gin"

type testHeader struct {
	Username string `header:"username"`
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

		c.JSON(200, gin.H{
			"message": "hello "+h.Username+"!! Login Successful!!",
		})
	})
	return r
}

func main() {
	RunServer().Run(":5000")
}

