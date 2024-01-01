package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// --------------------------
type UserInfo struct {
	username string
	password string
	role string
}
var USERS = map[string]UserInfo{
	"user1": UserInfo{
   username: "user1",
   password: "password1",
   role:     "admin",
 },
	"user2": UserInfo{
   username: "user2",
   password: "password2",
   role:     "member",
 },
}

// --------------------------
type Session struct {
	username string
	expiry   time.Time
}

var sessions = map[string]Session{}

func (s Session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

// --------------------------
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// --------------------------
func Login(w http.ResponseWriter, r *http.Request) {
	log.Print("DEBUG:Login")

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userInfo, ok := USERS[creds.Username]

	if !ok || userInfo.password != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(100 * time.Hour)

	sessions[sessionToken] = Session{
		username: creds.Username,
		expiry:   expiresAt,
	}

	http.SetCookie(w, &http.Cookie{
  Name:    "session_token",
  Value:   sessionToken,
		Expires: expiresAt,
  HttpOnly: true,
  Path : "/",
	})
}

// --------------------------
func Welcome(w http.ResponseWriter, r *http.Request) {
	log.Print("DEBUG:Welcome")

	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	userSession, exists := sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.username)))
}

// --------------------------
func Refresh(w http.ResponseWriter, r *http.Request) {
	log.Print("DEBUG:Refresh")

	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	userSession, exists := sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	sessions[newSessionToken] = Session{
		username: userSession.username,
		expiry:   expiresAt,
	}

	delete(sessions, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
}

// --------------------------
func Logout(w http.ResponseWriter, r *http.Request) {
	log.Print("DEBUG:Logout")

	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	delete(sessions, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
}
