// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Sex      int64  `json:"sex"`
	Email    string `json:"email"`
	Info     string `json:"info"`
}

type RegisterReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname,optional"`
	Sex      int64  `json:"sex,optional"`
	Email    string `json:"email,optional"`
	Info     string `json:"info,optional"`
}

type RegisterResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type LoginReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	User user `json:"userinfo"`
}

type UpdateUserInfoReq struct {
	User user `json:"userinfo"`
}

type UpdateUserInfoResp struct {
	UpdateInfo string `json:"result"`
}