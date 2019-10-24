package httphandler

import (
	"github.com/LuckyWindsck/sfc-course-guide/sfc-course-guide-server/internal/pkg/elasticclient"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "Hello, Gin",
		"message": "SFC COURSE GUIDE SERVER",
	})
}

func GetSearch(c *gin.Context) {
	query := c.Query("query")
	pretty := c.Query("pretty")
	var clientSearchResult elasticclient.ClientSearchResult
	var err error
	if query == "" {
		clientSearchResult, err = elasticclient.GetAllCourse()
	} else {
		clientSearchResult, err = elasticclient.SearchCourse(query)
	}

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	if pretty == "true" {
		c.IndentedJSON(http.StatusOK, clientSearchResult)
	} else {
		c.JSON(http.StatusOK, clientSearchResult)
	}
}
