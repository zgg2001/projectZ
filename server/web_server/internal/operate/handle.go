package operate

import (
	"net/http"

	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
)

func HandleRootRequest(w http.ResponseWriter, r *http.Request) {
	ok := checkSession(r)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func HandleLoginRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		uid, result := checkLogin(username, password)
		switch result {
		case rpc.LoginResult_LOGIN_SUCCESS:
			createSession(w, uid)
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		case rpc.LoginResult_LOGIN_FAIL_WRONG_PASSWORD:
			http.Error(w, "密码错误", http.StatusBadRequest)
		case rpc.LoginResult_LOGIN_FAIL_NOT_EXIST:
			http.Error(w, "用户不存在", http.StatusBadRequest)
		}
	}
}

func HandleLogoutRequest(w http.ResponseWriter, r *http.Request) {
	destroySession(w, r)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
