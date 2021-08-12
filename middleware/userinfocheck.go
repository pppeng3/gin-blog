package middleware

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tiia "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tiia/v20190529"
)

func CheckUserInfoMiddleware(c *gin.Context) {
	email := c.Request.FormValue("email")
	phonenumber := c.Request.FormValue("phone_number")
	avatarurl := c.Request.FormValue("avatar_url")
	//检验邮箱格式
	if email != "" {
		pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
		reg := regexp.MustCompile(pattern)
		if reg.MatchString(email) == false {
			c.JSON(http.StatusOK, gin.H{
				"code": 2009,
				"msg":  "请输入正确的邮箱地址",
			})
			c.Abort()
			return
		}
	}
	//检验电话号码格式
	if phonenumber != "" {
		regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
		pho := regexp.MustCompile(regular)
		if pho.MatchString(phonenumber) == false {
			c.JSON(http.StatusOK, gin.H{
				"code": 2010,
				"msg":  "请输入正确的电话号码",
			})
			c.Abort()
			return
		}
	}
	//检验用户头像
	if avatarurl != "" {
		credential := common.NewCredential(
			"AKIDvmU0VyPo5oSTO5ghhYHtPm0BHSZaPDh2", //SecretId
			"02OrCUginrnbIpmfS8fDewFVdcXM2gCf",     //SecretKey
		)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "tiia.tencentcloudapi.com"
		client, _ := tiia.NewClient(credential, "ap-guangzhou", cpf)

		request := tiia.NewDetectMisbehaviorRequest()

		request.ImageUrl = common.StringPtr(avatarurl)

		response, err := client.DetectMisbehavior(request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			logrus.Warnf("An API error has returned: %s", err)
			return
		}
		if err != nil {
			panic(err)
		}

		mini := float64(0.6)
		response.ToJsonString()
		if *response.Response.Confidence > mini {
			logrus.Info("picture is not be used")
			c.JSON(http.StatusOK, gin.H{
				"code": 2011,
				"msg":  "请放置符合要求的图片",
			})
			c.Abort()
			return
		}
		logrus.Infof("%s", response.ToJsonString())
	}
	c.Next()
}
