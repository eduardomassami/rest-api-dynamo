package domain

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	Save(user User) error
	FindByID(id string) (*User, error)
}
