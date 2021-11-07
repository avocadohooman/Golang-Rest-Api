package domain

type UserRepo interface {
	GetByEmail(email string) (*User, error)
	GetByUsername(name string) (*User, error)
	Create(user *User) (*User, error)
}

type DB struct {
	UserRepo UserRepo
}

type Domain struct {
	DB DB
}
