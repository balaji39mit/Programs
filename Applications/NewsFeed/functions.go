package main

import (
	"fmt"
)

var features map[string]func()

func init() {
	features = make(map[string]func(), 0)
	features["post"] = PostFeed
	features["follow"] = FollowUser
	features["comment"] = AddComment
	features["vote"] = Vote
	features["feed"] = ShowFeed
	features["comment-vote"] = VoteForComment
}

var (
	userNotExistError = "User does not exists."
)

var PostFeed = func() {
	fmt.Println("Enter userId and content of the post :")
	var userId int64
	var content string
	_, _ = fmt.Scanln(&userId, &content)
	createPost(content)
	if val, ok := Users[userId]; ok { //User exists.
		val.Posts = append(val.Posts, postCount)
	} else {
		Users[userId] = createUser(userId)
	}
}

var FollowUser = func() {
	fmt.Println("Enter the userId and the following userId : ")
	var userId, followingId int64
	_, _ = fmt.Scanln(&userId, &followingId)
	if userExists(userId) && userExists(followingId) {
		//update following and his posts
		updateFollow(userId, followingId)
	} else {
		fmt.Println(userNotExistError)
	}
}

var AddComment = func() {
	fmt.Println("Enter the userId, postId and the comment : ")
	var userId, postId int64
	var comment string
	_, _ = fmt.Scanln(&userId, &postId, &comment)
	if user, ok := Users[userId]; ok {
		if user.canComment(postId) {
			postComment(postId, comment)
		}
	} else {
		fmt.Println(userNotExistError)
	}

}

var Vote = func() {
	fmt.Println("Enter the userId, postId and the comment : ")
	var userId, postId int64
	_, _ = fmt.Scanln(&userId, &postId)

	if user, ok := Users[userId]; ok {
		if user.canComment(postId) {
			fmt.Println("Enter action : UPVOTE/DOWNVOTE")
			var action string
			_, _ = fmt.Scanln(&action)
			updateVote(postId, action)
		}
	} else {
		fmt.Println(userNotExistError)
	}
}

var VoteForComment = func() {
	fmt.Println("Enter the userId, postId and the comment Id : ")
	var userId, postId, commentId int64
	_, _ = fmt.Scanln(&userId, &postId, &commentId)
	if user, ok := Users[userId]; ok {
		if user.canComment(postId) {
			fmt.Println("Enter action : UPVOTE/DOWNVOTE")
			var action string
			_, _ = fmt.Scanln(&action)
			updateVoteForComments(postId, commentId, action)
		}
	} else {
		fmt.Println(userNotExistError)
	}
}

var ShowFeed = func() {
	fmt.Println("Enter the userId :")
	var userId int64
	_, _ = fmt.Scanln(&userId)
	if _, ok := Users[userId]; ok {
		feedPosts := getFeed(userId)
		for _, val := range feedPosts {
			printPost(val)
		}
	} else {
		fmt.Println(userNotExistError)
	}

}
