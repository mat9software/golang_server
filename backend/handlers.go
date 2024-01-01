package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// --------------------------
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
// --------------------------
// Parse session token from Cookie and populate http error code if applicable
func ParseSession(w *http.ResponseWriter, r *http.Request) (string, error) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			(*w).WriteHeader(http.StatusUnauthorized)
			return "", err
		}
		(*w).WriteHeader(http.StatusBadRequest)
		return "", err
	}
	sessionToken := c.Value
 return sessionToken, nil
}

// --------------------------
func Login(w http.ResponseWriter, r *http.Request) {
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

	SESSIONS[sessionToken] = Session{
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
 sessionToken, err := ParseSession(&w, r)
 if err != nil {
   return
 }

	userSession, exists := SESSIONS[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if userSession.isExpired() {
		delete(SESSIONS, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.username)))
}

// --------------------------
func Refresh(w http.ResponseWriter, r *http.Request) {
 sessionToken, err := ParseSession(&w, r)
 if err != nil {
   return
 }

	userSession, exists := SESSIONS[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if userSession.isExpired() {
		delete(SESSIONS, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	SESSIONS[newSessionToken] = Session{
		username: userSession.username,
		expiry:   expiresAt,
	}

	delete(SESSIONS, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
}

// --------------------------
func Logout(w http.ResponseWriter, r *http.Request) {
 sessionToken, err := ParseSession(&w, r)
 if err != nil {
   return
 }

	delete(SESSIONS, sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
}
