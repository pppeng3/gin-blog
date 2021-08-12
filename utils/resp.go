package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RetData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"msg": "success",
		"status_code": 0,
		"data": data,
	})
}

func RetErr(c *gin.Context, err error) {
	if err != nil {
		logrus.Error("Error:" +err.Error())
		c.JSON(200, gin.H{
			"msg": err,
			"status_code": -1,
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "success",
			"status_code": -1,
		})
	}
}

func RetSuccess(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "success",
		"status_code": 200,
	})
}
