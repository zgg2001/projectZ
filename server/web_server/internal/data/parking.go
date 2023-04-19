package data

import "github.com/zgg2001/projectZ/server/user_server/pkg/rpc"

type LoginRet struct {
	Uid int32           `json:"uid"`
	Ret rpc.LoginResult `json:"ret"`
}
