package main

import (
	"net/http"
)

// --------------------------------------------------
func UserHandlersInit() {
	http.HandleFunc("/users/get_all", UsersGetAllUsersAndRoles)
}

// --------------------------------------------------
// Define Public UserInfoPublic

// --------------------------------------------------
type UserRole int
const (
 USER_ROLE_ADMIN = iota
 USER_ROLE_MEMBER
 USER_ROLE_GUEST
)

// --------------------------------------------------
type UserInfo struct {
	username string
	password string
	role UserRole
}

// --------------------------------------------------
type UserInfoProtected struct {
	username string
	role UserRole
}


