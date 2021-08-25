package handler

import (
	"net/http"
	"strings"

	"github.com/cfabrica46/search-engine/database"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	var search struct {
		Sentence string `json:"sentence"`
	}

	err := c.BindJSON(&search)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ErrMessage": "Bad Request",
		})
		return
	}

	query := strings.ReplaceAll(search.Sentence, " ", ",")

	products, err := database.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	c.JSON(http.StatusOK, products)
}
