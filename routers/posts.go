package routers

import (
	controllers "github.com/cetRide/rideyu-api/controllers/posts"
	"github.com/gorilla/mux"
)

func GetPostRoutes(r *mux.Router) {
	// r.Use(middlewares.JwtAuthentication)
	r.HandleFunc("/post/create-post", controllers.CreatePost).Methods("POST")
	r.HandleFunc("/post/fetch-all-posts", controllers.FetchPosts).Methods("GET")
	r.HandleFunc("/post/fetch-user-posts/{userID}", controllers.FetchUserPosts).Methods("GET")
	r.HandleFunc("/post/fetch-a-post/{postID}", controllers.FetchSinglePost).Methods("GET")
	r.HandleFunc("/post/like/{category}/{id}", controllers.Like).Methods("POST")

	//comments
	r.HandleFunc("/post/create-comment/{post_id}", controllers.Comment).Methods("POST")

	//Replies
	r.HandleFunc("/comment/create-reply/{comment_id}", controllers.ReplyToComment).Methods("POST")
}
