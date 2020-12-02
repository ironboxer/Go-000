package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/lttzzlll/week02/service"
)

func main() {
	var s *service.TagService
	r := gin.Default()
	r.GET("/tags/:id", func(c *gin.Context) {
		id := c.Param("id")
		tagID, err := strconv.Atoi(id)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid Tag ID %s", id)
			return
		}

		tag, err := s.GetTag(uint64(tagID))
		if err != nil {
			c.String(http.StatusNotFound, "No Such Tag %s", id)
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": tag.ID, "name": tag.Name})
	})
	r.Run()
}
