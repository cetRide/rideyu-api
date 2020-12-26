package helpers

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func GetSession(request *http.Request) (userID uint64) {
	if cookie, err := request.Cookie("userID"); err == nil {
		cookieValue := make(map[string]uint64)
		if err = cookieHandler.Decode("userID", cookie.Value, &cookieValue); err == nil {
			userID = cookieValue["userID"]
		}
	}
	return userID
}

func DestroySession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "userID",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
func CreateSession(userID uint64, w http.ResponseWriter, r *http.Request) {

	value := map[string]uint64{
		"userID": userID,
	}
	if encoded, err := cookieHandler.Encode("userID", value); err == nil {
		cookie := &http.Cookie{
			Name:   "userID",
			Value:  encoded,
			Path:   "/",
			MaxAge: 3600,
		}
		http.SetCookie(w, cookie)
	}

}
