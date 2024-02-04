
package main

import (
	"net/http"
)

// --------------------------------------------------
func UsersGetAllUsersAndRoles(w http.ResponseWriter, r *http.Request) {
 err := SessionParseValidateRole(&w, r, USER_ROLE_ADMIN)
 if err != nil {
   return
 }

//todo
}
