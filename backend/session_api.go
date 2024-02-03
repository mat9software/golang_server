package main

import (
	"net/http"
)

// --------------------------------------------------
func SessionHandlersInit() {
	http.HandleFunc("/login", SessionLogin)
	http.HandleFunc("/welcome", SessionWelcome)
	http.HandleFunc("/refresh", SessionRefresh)
	http.HandleFunc("/logout", SessionLogout)
}

// --------------------------------------------------
type SessionCredentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
