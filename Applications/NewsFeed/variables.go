package main

var (
	Posts     map[int64]*Post
	Users     map[int64]*User
	postCount int64
)

func init() {
	postCount = 0
	Posts = make(map[int64]*Post, 0)
	Users = make(map[int64]*User, 0)
}
