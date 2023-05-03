package data

import "github.com/zgg2001/projectZ/server/user_server/pkg/rpc"

type RegisterRet struct {
	Ret rpc.RegistrationResult `json:"ret"`
}

type LoginRet struct {
	Uid int32           `json:"uid"`
	Ret rpc.LoginResult `json:"ret"`
}

type InfoRet struct {
	Cars []*CarInfo `json:"cars"`
}

type RechargeRet struct {
	Balance int32 `json:"balance"`
}

type OperatorRet struct {
	Ret rpc.CarOperationResult `json:"ret"`
}

type CarInfo struct {
	License      string `json:"license"`
	PTemperature int32  `json:"p_temperature"`
	PHumidity    int32  `json:"p_humidity"`
	PWeather     int32  `json:"p_weather"`
	PAddress     string `json:"p_address"`
	SID          int32  `json:"s_id"`
	STemperature int32  `json:"s_temperature"`
	SHumidity    int32  `json:"s_humidity"`
	SAlarm       int32  `json:"s_alarm"`
}
