package main

import (
	"time"
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
//TODO Add User validation here.

// --------------------------
type Session struct {
	username string
	expiry   time.Time
}

//TODO Add sessions validation here.
var SESSIONS = map[string]Session{}

func (s Session) isExpired() bool {
	return s.expiry.Before(time.Now())
}
