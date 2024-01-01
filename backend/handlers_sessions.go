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
func SessionLogin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userInfo, ok := UsersGetInfo(creds.Username)

	if !ok || userInfo.password != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(100 * time.Hour)

	SessionAdd(Session{
		username: creds.Username,
  token: newSessionToken,
		expiry:   expiresAt,
	})

	http.SetCookie(w, &http.Cookie{
  Name:    "session_token",
  Value:   newSessionToken,
		Expires: expiresAt,
  HttpOnly: true,
  Path : "/",
	})
}

// --------------------------
func SessionWelcome(w http.ResponseWriter, r *http.Request) {
 userSession, err := SessionParseValidate(&w, r)
 if err != nil {
   return
 }

	w.Write([]byte(fmt.Sprintf("Welcome %s!", userSession.username)))
}

// --------------------------
func SessionRefresh(w http.ResponseWriter, r *http.Request) {
 userSession, err := SessionParseValidate(&w, r)
 if err != nil {
   return
 }

	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(120 * time.Second)

	SessionAdd(Session{
		username: userSession.username,
  token: newSessionToken,
		expiry:   expiresAt,
	})

	SessionDelete(userSession)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})
}

// --------------------------
func SessionLogout(w http.ResponseWriter, r *http.Request) {
 sessionToken, err := SessionParseToken(&w, r)
 if err != nil {
   return
 }

	SessionDeleteToken(sessionToken)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
}
