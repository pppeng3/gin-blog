package rpc

import (
	user_center "Blog/proto/user-center"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserCenterServer struct {
}

func (UserCenterServer) Register(ctx context.Context, req *user_center.RegisterRequest) (*user_center.RegisterResponse, error) {
	return Register(ctx, req.User, req.UserExtra)
}
func (UserCenterServer) Login(ctx context.Context, req *user_center.LoginRequest) (*user_center.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UserCenterServer) Delete(ctx context.Context, req *user_center.DeleteRequest) (*user_center.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UserCenterServer) CheckToken(ctx context.Context, req *user_center.CheckTokenRequest) (*user_center.CheckTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckToken not implemented")
}
func (UserCenterServer) Refresh(ctx context.Context, req *user_center.RefreshRequest) (*user_center.RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UserCenterServer) GetUserInfo(ctx context.Context, req *user_center.GetUserInfoRequest) (*user_center.GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UserCenterServer) AddUser(ctx context.Context, req *user_center.AddUserRequest) (*user_center.AddUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
