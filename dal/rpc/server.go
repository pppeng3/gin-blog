package rpc

import (
	user_center "Blog/proto/user-center"
	"context"
)

type Server struct{}

func (server *Server) Register(ctx context.Context, in *user_center.RegisterRequest) (*user_center.RegisterResponse, error) {
	return nil, nil
}

func (server *Server) Login(ctx context.Context, in *user_center.LoginRequest) (*user_center.LoginResponse, error) {
	return nil, nil
}

func (server *Server) Delete(ctx context.Context, in *user_center.DeleteRequest) (*user_center.DeleteResponse, error) {
	return nil, nil
}

func (server *Server) GetUserInfo(ctx context.Context, in *user_center.GetUserInfoRequest) (*user_center.GetUserInfoResponse, error) {
	return nil, nil
}

func (server *Server) CheckToken(ctx context.Context, in *user_center.CheckTokenRequest) (*user_center.CheckTokenResponse, error) {
	return nil, nil
}

func (server *Server) Refresh(ctx context.Context, in *user_center.RefreshRequest) (*user_center.RefreshResponse, error) {
	return nil, nil
}

func (server *Server) AddUser(ctx context.Context, in *user_center.AddUserRequest) (*user_center.AddUserResponse, error) {
	return nil, nil
}

func main() {
	// lis, err := net.Listen("tcp", "3333")
	// if err != nil {
	// 	logrus.Fatalf("failed to listen : %v", err)
	// }
	// ser := grpc.NewServer()
	// user_center.RegisterUserCenterServer(ser, &Server{})
	// reflection.Register(ser)
	// if err := ser.Serve(lis); err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
}
