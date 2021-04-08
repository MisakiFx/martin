package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPageSizeFromQuery(c *gin.Context) (int, int, error) {
	pageString := c.DefaultQuery("page", "0")
	sizeString := c.DefaultQuery("size", "0")
	page, err := strconv.ParseInt(pageString, 10, 32)
	if err != nil {
		return 0, 0, err
	}
	size, err := strconv.ParseInt(sizeString, 10, 32)
	if err != nil {
		return 0, 0, err
	}
	return int(page), int(size), nil
}
