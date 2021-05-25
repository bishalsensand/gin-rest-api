package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/entry", getEntries)
	r.POST("/entry", createEntry)
	r.GET("/entry/:id", getEntry)
	r.PUT("/entry/:id", putEntry)
	r.DELETE("/entry/:id", deleteEntry)

	r.Run()
}
