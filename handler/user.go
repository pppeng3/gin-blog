package handler

import (
	"Blog/dal/rpc"
	"Blog/model"
	"Blog/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Register(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.RegisterReq
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
		return
	}
	resp, err := rpc.Register(ctx, &reqBody.User, &reqBody.Userextra)
	if err != nil {
		logrus.Warnf("register fail: %v", err)
		utils.RetErr(c, err)
		return
	}
	var RegisterRes model.RegisterResp
	RegisterRes.RegisterResponse = resp
	utils.RetData(c, RegisterRes)
}

func Login(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.LoginReq
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
		return
	}
	logrus.Infof("reqBody: %+v", reqBody)
	resp, err := rpc.Login(ctx, reqBody.Username, reqBody.Password)
	if err != nil {
		logrus.Warnf("login failed, err: %v", err)
		utils.RetErr(c, err)
		return
	}
	var loginRes model.LoginResp
	loginRes.LoginResponse = resp
	loginRes.Auth = resp.AuthList
	utils.RetData(c, loginRes)
}

func Delete(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.DeleteReq
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
		return
	}
	logrus.Infof("delete: %+v", reqBody)
	resp, err := rpc.Delete(ctx, reqBody.Id)
	if err != nil {
		logrus.Warnf("delete failed, err : %v", err)
		utils.RetErr(c, err)
		return
	}
	var deleteResp model.DeleteResp
	deleteResp.DeleteResponse = resp
	utils.RetData(c, deleteResp)
}

func GetUserInfo(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.GetUserInfoReq
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
		return
	}
	resp, err := rpc.GetUserInfo(ctx, reqBody.Username) //删掉id
	if err != nil {
		logrus.Warnf("get userinfo failed, err : %v", err)
		utils.RetErr(c, err)
		return
	}
	var getResp model.GetUserInfoResp
	getResp.GetUserInfoResponse = resp
	utils.RetData(c, getResp)
}
