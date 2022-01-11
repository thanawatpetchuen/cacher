package main

import "errors"

type UserRepository interface {
	GetOne(id int) (User, error)
	GetMany(id ...int) ([]User, error)
}

type userRepository struct {
}

func (ur userRepository) GetOne(id int) (User, error) {
	for _, u := range usersDB {
		if u.Id == id {
			return u, nil
		}
	}
	return User{}, errors.New("not found")
}

func (ur userRepository) GetMany(id ...int) ([]User, error) {
	return []User{}, nil
}

func NewUserRepository() UserRepository {
	return userRepository{}
}
