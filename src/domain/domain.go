package domain

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	Save(user User) error
	FindByID(id string) (*User, error)
}

type UserBucket interface {
	Save(user []byte) error
}
