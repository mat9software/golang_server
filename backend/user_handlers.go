
package main

import (
 "encoding/json"
	"net/http"
)

// --------------------------------------------------
func UsersGetAllUsersAndRoles(w http.ResponseWriter, r *http.Request) {
//mdtmp skip role for testing
/*
 err := SessionParseValidateRole(&w, r, USER_ROLE_ADMIN)
 if err != nil {
   return
 }
*/

 allUsers := UsersGetAllProtected();

 jsonData, err := json.Marshal(allUsers)//mdtmp this doesn't work.
 if err != nil {
  http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
  return
 }

 w.Header().Set("Content-Type", "application/json")

 w.Write(jsonData)
}
