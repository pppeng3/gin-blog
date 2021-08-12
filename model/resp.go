package model

import user_center "Blog/proto/user-center"

type RegisterResp struct {
	*user_center.RegisterResponse
}

type LoginResp struct {
	*user_center.LoginResponse
	Auth []string `json:"auth"`
}

type DeleteResp struct {
	*user_center.DeleteResponse
}

type GetUserInfoResp struct {
	*user_center.GetUserInfoResponse
}
