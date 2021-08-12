package middleware

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

const (
	levelD = iota
	levelC
	levelB
	levelA
	levelS
)

//密码长度
const (
	minLength = 6
	maxLength = 12
)

func CheckRegisterMiddleware(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	if len(username) < 3 {
		c.JSON(http.StatusOK, gin.H{
			"code": 2006,
			"msg":  "用户名长度不能小于3",
		})
		c.Abort()
		return
	}
	if len(password) < minLength {
		c.JSON(http.StatusOK, gin.H{
			"code": 2007,
			"msg":  "密码设置长度不能小于6",
		})
		c.Abort()
		return
	}
	if len(password) > maxLength {
		c.JSON(http.StatusOK, gin.H{
			"code": 2008,
			"msg":  "密码设置长度不能大于12",
		})
		c.Abort()
		return
	}
	var level int = levelD
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, password)
		if match {
			level++
		}
	}
	if level < levelA {
		c.JSON(http.StatusOK, gin.H{
			"code": 2009,
			"msg":  "账户密码安全性太低,必须包含数字、大写字母、小写字母、标点符号中的三种",
		})
		c.Abort()
		return
	}
	c.Next()
}
