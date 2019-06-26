package user

import "time"

type User struct {
	Id        int64
	Name      string
	CreatedOn time.Time
	UpdatedOn time.Time
}

func (u User) String() string {
	return u.Name
}
