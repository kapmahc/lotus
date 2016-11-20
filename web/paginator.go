package web

import (
	"math"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

//ParsePager parse page and size
func ParsePager(c *gin.Context, s int) (int, int, int) {
	size, err := strconv.ParseInt(c.Query("size"), 10, 32)
	if err != nil {
		log.Error(err)
	}
	page, err := strconv.ParseInt(c.Query("page"), 10, 32)
	if page <= 0 {
		page = 1
	}
	return int(page), int(size), int((page - 1) * size)
}

//Paginator paginator
func Paginator(current, size int, count int64) gin.H {

	total := int(math.Ceil(float64(count) / float64(size)))
	if current > total {
		current = total
	}
	if current <= 0 {
		current = 1
	}

	var prev int
	prev = current - 1
	if prev <= 0 {
		prev = 1
	}

	var next int
	next = current + 1
	if next > total {
		next = total
	}

	return gin.H{
		"total":   total,
		"prev":    prev,
		"next":    next,
		"size":    size,
		"current": current,
	}
}
