package user

var (
	Client IUser
	users  []*User
)

type Impl struct {
}

func init() {
	Client = Impl{}
	users = []*User{
		{
			Id:   1,
			Name: "Balaji",
		},
		{
			Id:   2,
			Name: "test",
		},
	}
}
func (u Impl) Get(id int64) *User {
	return u.getUser(id)
}

func (u Impl) getUser(id int64) *User {
	for _, val := range users {
		if val.Id == id {
			return val
		}
	}
	return nil
}

func (u Impl) Display(id int64) string {
	user := u.getUser(id)
	if user != nil {
		return user.Name
	}
	return ""
}
