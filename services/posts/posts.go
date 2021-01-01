package posts

import (
	"strconv"
	"strings"

	"github.com/cetRide/rideyu-api/apihelpers"
	"github.com/cetRide/rideyu-api/models"
)

func CreatePost(data map[string]interface{}, userID uint64) map[string]interface{} {

	description := data["description"].(string)

	post := models.Post{Description: description, UserID: userID}

	files := data["files"].([]string)

	tx := models.GetDB().Begin()

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&post).
		Error; err != nil {
		tx.Rollback()
		return apihelpers.Message(false, "Unable to create a post")
	}

	for i := range files {
		if err := tx.Create(&models.PostMedia{
			PostID:    post.ID,
			MediaType: "post",
			FileName:  files[i]}).
			Error; err != nil {
			tx.Rollback()
			return apihelpers.Message(false, "Unable to create a post")
		}
	}

	tx.Commit()
	return apihelpers.Message(true, "Post Created Successfully")
}

func GetAllPosts() map[string]interface{} {
	var posts []models.PostMedia
	data := models.GetDB().Find(&posts)
	return map[string]interface{}{"status": true, "data": data}
}

func FetchUserPosts(userID string) map[string]interface{} {
	user, _ := strconv.ParseUint(userID, 10, 64)
	var posts []models.Post
	data := models.GetDB().Where("user_id = ?", user).Find(&posts)
	return map[string]interface{}{"status": true, "data": data}
}

func FetchSinglePost(postID string) map[string]interface{} {
	postId, _ := strconv.ParseUint(postID, 10, 64)
	var post models.Post
	data := models.GetDB().Find(&post, postId)
	return map[string]interface{}{"status": true, "data": data}
}

func Like(item_id string, category string, userID uint64) map[string]interface{} {

	id, _ := strconv.ParseUint(item_id, 10, 64)
	if err := models.GetDB().Create(&models.Like{
		UserID:   userID,
		ItemID:   id,
		Category: category}).
		Error; err != nil {
		return apihelpers.Message(false, "Unable to like")
	}
	return apihelpers.Message(true, strings.Title(category)+" liked")
}

func CreateComment(user uint64, postID string, comment string) map[string]interface{} {
	post_iD, _ := strconv.ParseUint(postID, 10, 64)
	if err := models.GetDB().Create(&models.Comment{
		PostID:  post_iD,
		Comment: comment,
		UserID:  user,
	}).Error; err != nil {
		return apihelpers.Message(false, "Unable to comment")
	}
	return apihelpers.Message(true, "Comment sent!")
}

func CreateReply(user uint64, commentID string, reply string) map[string]interface{} {
	comment_iD, _ := strconv.ParseUint(commentID, 10, 64)
	if err := models.GetDB().Create(&models.CommentReply{
		CommentID: comment_iD,
		Reply:     reply,
		UserID:    user,
	}).Error; err != nil {
		return apihelpers.Message(false, "Unable to reply")
	}
	return apihelpers.Message(true, "Reply sent!")
}
