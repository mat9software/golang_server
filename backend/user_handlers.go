
package main

import (
 "encoding/json"
	"net/http"
"fmt"
)

type MyData struct {
	Role int
 Message string
 Message2 string
}

// --------------------------------------------------
func UsersGetAllUsersAndRoles(w http.ResponseWriter, r *http.Request) {
 err := SessionParseValidateRole(&w, r, USER_ROLE_ADMIN)
 if err != nil {
   return
 }

 allUsers := UsersGetAllProtected()

 fmt.Print(allUsers)

 jsonData, err := json.Marshal(allUsers)
 fmt.Print(string(jsonData))
 if err != nil {
  http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
  return
 }

 w.Header().Set("Content-Type", "application/json")

 w.Write(jsonData)
}
