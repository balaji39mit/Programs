package main

import (
	"fmt"
	"time"
)

type User struct {
	Id        int64
	Following []int64
	Followers []int64
	Posts     []int64
}

func (user *User) AddFollow(followingId int64) {
	user.Following = append(user.Following, followingId)
	followingUser, _ := Users[followingId]
	//check whether user should be able to see transitive dependency posts
	//user.Posts = append(user.Posts, followingUser.Posts...)
	//update followers.
	followingUser.Followers = append(followingUser.Followers, user.Id)
}

func (user *User) canComment(postId int64) bool {
	if containsInt64(getFeed(user.Id), postId) {
		return true
	} else {
		fmt.Println("This user cannot comment on this post.")
	}
	return false
}

type Post struct {
	Id        int64
	Content   string
	Comments  []*Comment
	UpVotes   int64
	DownVotes int64
	CreatedOn time.Time
}

func (post *Post) UpdateVote(action string) {
	switch action {
	case upvote:
		post.UpVotes = post.UpVotes + 1
	case downvote:
		post.DownVotes = post.DownVotes + 1
	}
}

func (post *Post) AddComment(comment string) {
	cmt := Comment{
		Id:      int64(len(post.Comments) + 1),
		Content: comment,
	}
	post.Comments = append(post.Comments, &cmt)
}

func (post *Post) canVote(commentId int64) bool {
	if containsComment(post.Comments, commentId) {
		return true
	} else {
		fmt.Println("This user cannot comment on this post.")
	}
	return false
}

type Comment struct {
	Id        int64
	Content   string
	UpVotes   int64
	DownVotes int64
	Reply     *Comment
}

func (cmt *Comment) UpdateVote(action string) {
	switch action {
	case upvote:
		cmt.UpVotes = cmt.UpVotes + 1
	case downvote:
		cmt.DownVotes = cmt.DownVotes + 1
	}
}
