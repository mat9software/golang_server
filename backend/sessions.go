package main

import (
	"time"
 "errors"
	"net/http"
)

// --------------------------
// Errors
var(
SessionNotFoundErr = errors.New("Session not found")
SessionExpiredErr = errors.New("Session expired")
)

// --------------------------
type UserInfo struct {
	username string
	password string
	role string
}
type UserMap map[string]UserInfo

var USERS = UserMap {
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

func UsersGetInfo(username string) (UserInfo, bool) {
	 userInfo, ok := USERS[username]
  return userInfo, ok
}

// --------------------------
type Session struct {
	username string
	token string
	expiry   time.Time
}

var SESSIONS = map[string]Session{}

func (s Session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func SessionAdd(session Session) {
	SESSIONS[session.token] = session
}
func SessionDelete(session Session) {
	delete(SESSIONS, session.token)
}
func SessionDeleteToken(sessionToken string) {
	delete(SESSIONS, sessionToken)
}

// --------------------------
// Parse and validate session token. Populate http error code if applicable
func SessionParseValidate(w *http.ResponseWriter, r *http.Request)(Session, error) {
 sessionToken, err := SessionParseToken(w, r)
 if err != nil {
   return Session{}, err 
 }

 userSession, err := SessionValidateToken(w, r, sessionToken)
 if err != nil {
   return userSession, err
 }

 return userSession, err
}

// --------------------------
// Parse session token from Cookie and populate http error code if applicable
func SessionParseToken(w *http.ResponseWriter, r *http.Request) (string, error) {
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
// Validate if seesion token is active (user logged in). Populate http error code if applicable
func SessionValidateToken(w *http.ResponseWriter, r *http.Request, sessionToken string) (Session, error) {

	userSession, exists := SESSIONS[sessionToken]
	if !exists {
		(*w).WriteHeader(http.StatusUnauthorized)
		return userSession, SessionNotFoundErr
	}

	if userSession.isExpired() {
		delete(SESSIONS, sessionToken)
		(*w).WriteHeader(http.StatusUnauthorized)
		return userSession, SessionExpiredErr
	}

 return userSession, nil
}
