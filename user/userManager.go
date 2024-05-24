package user

// UserManager defines the methods for user management
type UserManager interface {
	AddUser(name, gender string, age int)
	GetUser(name string) (User, error)
	GetAllUsers() []User
}
