package rpc

import (
	user_center "Blog/proto/user-center"
	"context"
	"testing"
)

func TestLogin(t *testing.T) {
	_, err := userCenterClient.Login(context.Background(), &user_center.LoginRequest{})
	if err != nil {
		panic(err)
	}
}
