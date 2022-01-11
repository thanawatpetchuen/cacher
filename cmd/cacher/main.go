package main

import (
	"log"

	"github.com/thanawatpetchuen/cacher"
)

func main() {
	cacher := cacher.New()

	log.Println("Cacher", cacher)

	userRepository := NewUserRepository()
	userService := NewUserService(cacher, userRepository)

	user, err := userService.GetOne(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User:", user)

	user, err = userService.GetOne(2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User:", user)

	user, err = userService.GetOne(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("User:", user)

}
