package posts

import (
	"encoding/json"
	"net/http"

	"github.com/cetRide/rideyu-api/apihelpers"
	"github.com/cetRide/rideyu-api/helpers"
	post "github.com/cetRide/rideyu-api/services/posts"
	"github.com/gorilla/mux"
)

type RequestData struct {
	Comment string `json:"comment"`
	Reply   string `json:"reply"`
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSession(r)
	postData := helpers.UploadFiles(w, r, user, "post")
	response := post.CreatePost(postData, user)
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

func Comment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	comment := &RequestData{}
	post_id := params["post_id"]
	err := json.NewDecoder(r.Body).Decode(comment)
	if err != nil {
		apihelpers.Respond(w, apihelpers.Message(false, "Invalid request"))
		return
	}
	response := post.CreateComment(helpers.GetSession(r), post_id, comment.Comment)
	apihelpers.Respond(w, response)
}

func ReplyToComment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reply := &RequestData{}
	comment_id := params["comment_id"]
	err := json.NewDecoder(r.Body).Decode(reply)
	if err != nil {
		apihelpers.Respond(w, apihelpers.Message(false, "Invalid request"))
		return
	}
	response := post.CreateReply(helpers.GetSession(r), comment_id, reply.Reply)
	apihelpers.Respond(w, response)
}
