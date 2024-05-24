package user

import (
	"errors"
)

// UserManagerImpl implements the UserManager interface
type userManagerImpl struct {
	Users map[string]User
}

func NewUserManager() UserManager {
	return &userManagerImpl{Users: make(map[string]User)}
}

func (um *userManagerImpl) AddUser(name, gender string, age int) {
	user := User{Name: name, Gender: gender, Age: age}
	um.Users[name] = user
}

func (um *userManagerImpl) GetUser(name string) (User, error) {
	user, exists := um.Users[name]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func (um *userManagerImpl) GetAllUsers() []User {
	users := make([]User, 0, len(um.Users))
	for _, user := range um.Users {
		users = append(users, user)
	}
	return users
}
