package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lttzzlll/week02/dao"
	"log"
	"net/http"
	"strconv"

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
			var errMsg = err.Error()
			var returnCode int
			if dao.IsDBConnectError(err) {
				returnCode = http.StatusServiceUnavailable
			} else if dao.IsEmptyQueryResult(err) {
				returnCode = http.StatusNotFound
			} else {
				returnCode = http.StatusInternalServerError
			}
			c.String(returnCode, errMsg)
			log.Printf("Error: %s, Status: %d", errMsg, returnCode)
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": tag.ID, "name": tag.Name})
	})
	r.Run()
}
