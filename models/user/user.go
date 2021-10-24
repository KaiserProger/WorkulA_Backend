package user

type User struct {
	UserId   int
	Name     string
	Email    string
	Password string
}

func New(user_id int, name string, email string, password [32]byte) *User {
	return &User{
		UserId:   user_id,
		Name:     name,
		Email:    email,
		Password: string(password[:]),
	}
}
