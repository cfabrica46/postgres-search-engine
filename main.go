package main

import (
	"github.com/cfabrica46/search-engine/handler"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./files", false)))

	s := r.Group("/api/v1")
	{
		s.POST("/search", handler.Search)
	}
	r.Run(":8080")
}
