package posts

import (
	"strconv"
	"strings"

	"github.com/cetRide/rideyu-api/apihelpers"
	"github.com/cetRide/rideyu-api/models"
)

func CreatePost(description string, userID uint64) map[string]interface{} {

	if err := models.GetDB().Create(&models.Post{
		Description: description,
		UserID:      userID}).
		Error; err != nil {
		return apihelpers.Message(false, "Unable to create a post")
	}
	return apihelpers.Message(true, "Post Created Successfully")
}

func GetAllPosts() map[string]interface{} {
	var posts []models.Post
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
