package users

type User struct {
	Name string
	ID   string
}

func (u User) GetUserDetails() string {
	return "Name: " + u.Name + ", ID: " + u.ID
}
