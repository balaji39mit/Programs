package user

type IUser interface {
	Get(id int64) *User
	Display(id int64) string
}
