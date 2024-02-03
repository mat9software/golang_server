package main

type UserMap map[string]UserInfo

// --------------------------------------------------
var USERS = UserMap {
	"user1": UserInfo{
   username: "user1",
   password: "password1",
   role:     USER_ROLE_ADMIN,
 },
	"user2": UserInfo{
   username: "user2",
   password: "password2",
   role:     USER_ROLE_MEMBER,
 },
}

// --------------------------------------------------
func UsersGetInfo(username string) (UserInfo, bool) {
	 userInfo, ok := USERS[username]
  return userInfo, ok
}

func UsersValidatePassword(username string, password string) (bool) {
	userInfo, ok := UsersGetInfo(username)

	if !ok || userInfo.password != password {
		return false
	}
 return true
}

//TODO Use it
func UsersValidateRole(username string, minimumRole UserRole) (bool) {
	userInfo, ok := UsersGetInfo(username)

	if !ok || userInfo.role > minimumRole {
		return false
	}
 return true
}

