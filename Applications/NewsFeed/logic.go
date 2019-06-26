package main

import (
	"sort"
)

//getFeed should return the posts of this user
//and the post of the followers of current user

var SortLogic = func(post1, post2 int64) bool {
	a, _ := Posts[post1]
	b, _ := Posts[post2]

	aScore := a.UpVotes - a.DownVotes
	bScore := b.UpVotes - b.DownVotes
	if aScore > bScore {
		return true
	}
	if aScore < bScore {
		return false
	}
	if len(a.Comments) > len(b.Comments) {
		return true
	}
	if len(a.Comments) < len(b.Comments) {
		return false
	}
	if a.CreatedOn.After(b.CreatedOn) {
		return true
	}
	return false
}

var getFeed = func(userId int64) []int64 {
	user, _ := Users[userId]
	posts := user.Posts
	for _, val := range user.Following {
		following, _ := Users[val]
		posts = append(posts, following.Posts...)
	}
	//fmt.Print(len(posts))
	sort.Slice(posts, func(i, j int) bool {
		return SortLogic(posts[i], posts[j])
	})
	return posts
}
