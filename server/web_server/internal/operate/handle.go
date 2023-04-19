package operate

import (
	"encoding/json"
	"net/http"

	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
	"github.com/zgg2001/projectZ/server/web_server/internal/data"
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
	destroySession(w, r)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
