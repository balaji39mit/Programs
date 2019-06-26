package main

import (
	"fmt"
)

var userExists = func(userId int64) bool {
	_, ok := Users[userId]
	return ok
}

var containsInt64 = func(arr []int64, x int64) bool {
	for _, val := range arr {
		if x == val {
			return true
		}
	}
	return false
}

var containsComment = func(arr []*Comment, commentId int64) bool {
	for _, val := range arr {
		if val != nil && val.Id == commentId {
			return true
		}
	}
	return false
}

var printPost = func(postId int64) {
	post := Posts[postId]
	fmt.Println(fmt.Sprintf("Content  : %s, Upvotes : %d, DownVotes : %d, Comments : %v, len : %d", post.Content, post.UpVotes, post.DownVotes, post.Comments, len(post.Comments)))
}
