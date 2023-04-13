package operate

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
)

const (
	sessionDuration = time.Hour * 24
)

type Session struct {
	ID        string
	UserID    int32
	ExpiresAt time.Time
}

func createSession(w http.ResponseWriter, userID int32) *Session {
	b := make([]byte, 32)
	rand.Read(b)
	id := base64.URLEncoding.EncodeToString(b)
	expiresAt := time.Now().Add(sessionDuration)
	session := &Session{
		ID:        id,
		UserID:    userID,
		ExpiresAt: expiresAt,
	}
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    id,
		Expires:  expiresAt,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	return session
}

func checkSession(r *http.Request) bool {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return false
	}
	session := getSessionFromDB(cookie.Value)
	if session == nil || session.ExpiresAt.Before(time.Now()) {
		return false
	}
	return true
}

// 从数据库或缓存中获取session
func getSessionFromDB(id string) *Session {
	// TODO: 根据session ID查询数据库或缓存，返回session对象
	return nil
}

func destroySession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return
	}
	deleteSessionFromDB(cookie.Value)
	cookie.Expires = time.Now().Add(-time.Hour)
	http.SetCookie(w, cookie)
}

// 从数据库或缓存中删除session
func deleteSessionFromDB(id string) {
	// TODO: 根据session ID删除数据库或缓存中的session
}
