package operate

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
	"github.com/zgg2001/projectZ/server/web_server/internal/data"
)

func HandleRegisterRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		result := checkRegister(username, password)
		p := data.RegisterRet{Ret: result}
		jsonBytes, err := json.Marshal(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	}
}

func HandleLoginRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		uid, result := checkLogin(username, password)
		p := data.LoginRet{Uid: uid, Ret: result}
		if result == rpc.LoginResult_LOGIN_SUCCESS {
			createSession(w, uid)
		}
		jsonBytes, err := json.Marshal(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	}
}

func HandleLogoutRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		destroySession(w, r)
	}
}

func HandleInfoRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		uid, ok := checkSession(r)
		if !ok {
			http.Error(w, "未经授权", http.StatusUnauthorized)
			return
		}
		infos, err := getInfo(uid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var ret data.InfoRet
		for _, info := range infos {
			d := &data.CarInfo{
				License:      info.GetLicense(),
				PTemperature: info.GetPTemperature(),
				PHumidity:    info.GetPHumidity(),
				PWeather:     info.GetPWeather(),
				PAddress:     info.GetPAddress(),
				SID:          info.GetSId(),
				STemperature: info.GetSTemperature(),
				SHumidity:    info.GetSHumidity(),
				SAlarm:       int32(info.GetSAlarm()),
			}
			ret.Cars = append(ret.Cars, d)
		}
		jsonBytes, err := json.Marshal(ret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	}
}

func HandleRechargeRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		uid, ok := checkSession(r)
		if !ok {
			http.Error(w, "未经授权", http.StatusUnauthorized)
			return
		}
		amount, _ := strconv.Atoi(r.FormValue("amount"))
		balance := recharge(uid, int32(amount))
		p := data.RechargeRet{Balance: balance}
		jsonBytes, err := json.Marshal(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	}
}

func HandleOperatorRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		uid, ok := checkSession(r)
		if !ok {
			http.Error(w, "未经授权", http.StatusUnauthorized)
			return
		}
		operation := rpc.CarOperation_OPERATION_ADD
		license := r.FormValue("license")
		newLicense := ""
		switch r.FormValue("operation") {
		case "add":
			operation = rpc.CarOperation_OPERATION_ADD
		case "delete":
			operation = rpc.CarOperation_OPERATION_DELETE
		case "change":
			operation = rpc.CarOperation_OPERATION_CHANGE
			newLicense = r.FormValue("newLicense")
		}
		result := carOperator(uid, operation, license, newLicense)
		p := data.OperatorRet{Ret: result}
		jsonBytes, err := json.Marshal(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	}
}
