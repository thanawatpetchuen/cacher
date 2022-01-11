package main

import (
	"strconv"

	"github.com/thanawatpetchuen/cacher"
)

type User struct {
	Id   int
	Name string
}

type UserService interface {
	GetOne(id int) (User, error)
	GetMany(id ...int) ([]User, error)
}

type userService struct {
	cacher   cacher.Cacher
	userRepo UserRepository
}

func (us userService) GetOne(id int) (User, error) {
	user, err := us.userRepo.GetOne(id)
	if err != nil {
		return User{}, err
	}
	_, err = us.cacher.Do("GetOne"+strconv.Itoa(id), user)
	if err != nil {
		return User{}, err
	}
	return user, nil

}
func (us userService) GetMany(id ...int) ([]User, error) {
	return []User{{
		Id:   1,
		Name: "a",
	}}, nil
}

func NewUserService(cacher cacher.Cacher, userRepo UserRepository) UserService {
	return userService{cacher, userRepo}
}
