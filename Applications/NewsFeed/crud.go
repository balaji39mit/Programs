package main

import "time"

var createPost = func(content string) {
	postCount++
	Posts[postCount] = &Post{
		Id:        postCount,
		Content:   content,
		CreatedOn: time.Now(),
	}
}

var createUser = func(userId int64) *User {
	user := User{
		Id:    userId,
		Posts: []int64{postCount},
	}
	return &user
}

var updateVote = func(postId int64, action string) {
	post, _ := Posts[postId]
	post.UpdateVote(action)
}

var updateVoteForComments = func(postId int64, commentId int64, action string) {
	post, _ := Posts[postId]
	if post.canVote(commentId) {
		post.Comments[commentId].UpdateVote(action)
	}
}

var postComment = func(postId int64, comment string) {
	post := Posts[postId]
	post.AddComment(comment)
}

var updateFollow = func(userId, followingId int64) {
	user, _ := Users[userId]
	user.AddFollow(followingId)
}
