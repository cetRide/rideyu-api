package posts

import (
	"encoding/json"
	"net/http"

	"github.com/cetRide/rideyu-api/apihelpers"
	"github.com/cetRide/rideyu-api/helpers"
	post "github.com/cetRide/rideyu-api/services/posts"
	"github.com/gorilla/mux"
)

type PostData struct {
	Description string `json:"description"`
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	data := &PostData{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		apihelpers.Respond(w, apihelpers.Message(false, "Invalid request"))
		return
	}

	response := post.CreatePost(data.Description, helpers.GetSession(r))
	apihelpers.Respond(w, response)
}

func FetchPosts(w http.ResponseWriter, r *http.Request) {
	response := post.GetAllPosts()
	apihelpers.Respond(w, response)
}

func FetchUserPosts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["userID"]
	response := post.FetchUserPosts(userID)
	apihelpers.Respond(w, response)
}

func FetchSinglePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID := params["postID"]
	response := post.FetchSinglePost(postID)
	apihelpers.Respond(w, response)
}

func Like(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	category := params["category"]
	response := post.Like(id, category, helpers.GetSession(r))
	apihelpers.Respond(w, response)
}
