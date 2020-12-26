package users

import (
	"encoding/json"
	"net/http"

	u "github.com/cetRide/rideyu-api/apihelpers"
	data "github.com/cetRide/rideyu-api/helpers"
	models "github.com/cetRide/rideyu-api/services/users"
)

func Register(w http.ResponseWriter, r *http.Request) {
	user := &data.AccountDetails{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	response := models.CreateAccount(user.Username, user.Email, user.Password)
	u.Respond(w, response)
}

func Login(w http.ResponseWriter, r *http.Request) {

	user := &data.AccountDetails{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	response := models.LoginUser(user.Username, user.Password)

	if response["status"].(bool) {
		data.CreateSession(response["user"].(uint64), w, r)
	}
	u.Respond(w, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	data.DestroySession(w)
	u.Respond(w, u.Message(true, "Successfully logged out."))
}
