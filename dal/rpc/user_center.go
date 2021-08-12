package rpc

import (
	user_center "Blog/proto/user-center"
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func Register(ctx context.Context, user *user_center.User, extra *user_center.UserExtra) (resp *user_center.RegisterResponse, err error) {
	req := &user_center.RegisterRequest{
		User:      user,
		UserExtra: extra,
	}
	rep, err := userCenterClient.Register(ctx, req)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	logrus.Infof("注册成功, id: %d\n", rep.Id)
	return
}

func Login(ctx context.Context, username, password string) (resp *user_center.LoginResponse, err error) {
	req := &user_center.LoginRequest{
		Username: username,
		Password: password,
	}
	resp, err = userCenterClient.Login(ctx, req)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	return
}

func Delete(ctx context.Context, id int32) (resp *user_center.DeleteResponse, err error) {
	req := &user_center.DeleteRequest{
		Id: id,
	}
	resp, err = userCenterClient.Delete(ctx, req)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	return
}

func CheckToken(ctx context.Context, token string) (userInfo *user_center.UserInfo, err error) {
	resp, err := userCenterClient.CheckToken(ctx, &user_center.CheckTokenRequest{
		Token:     token,
		IsRefresh: false,
	})
	if err != nil {
		logrus.Warnln(err)
		return nil, errors.WithStack(err)
	}
	return resp.User, nil
}

func GetUserInfo(ctx context.Context, name string) (resp *user_center.GetUserInfoResponse, err error) {
	req := &user_center.GetUserInfoRequest{
		Username: name,
	}
	resp, err = userCenterClient.GetUserInfo(ctx, req)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	return
}

func AddUser(ctx context.Context, name string) (resp *user_center.AddUserResponse, err error) {
	req := &user_center.AddUserRequest{
		Username: name,
	}
	resp, err = userCenterClient.AddUser(ctx, req)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	return
}
