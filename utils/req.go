package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetParamsInt64(c *gin.Context, key string) int64 {
	res, err := strconv.ParseInt(c.DefaultQuery(key, "-1"), 10, 64)
	if err != nil {
		return -1
	}
	return res
}

func GetParamsStr(c *gin.Context, key string) string {
	return c.DefaultQuery(key, "")
}
