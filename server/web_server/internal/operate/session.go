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

var (
	sessionMap = map[string]*Session{}
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
	createSessionFromDB(id, session)
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    id,
		Expires:  expiresAt,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	return session
}

func checkSession(r *http.Request) (int32, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return -1, false
	}
	session := getSessionFromDB(cookie.Value)
	if session == nil || session.ExpiresAt.Before(time.Now()) {
		return -1, false
	}
	return session.UserID, true
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

func createSessionFromDB(id string, s *Session) {
	sessionMap[id] = s
}

func getSessionFromDB(id string) *Session {
	s, ok := sessionMap[id]
	if ok {
		return s
	}
	return nil
}

func deleteSessionFromDB(id string) {
	_, ok := sessionMap[id]
	if ok {
		delete(sessionMap, id)
	}
}
