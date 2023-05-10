package integration

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/services"
	"testing"
)

func TestCreateUser(t *testing.T) {
	userService := services.NewUserService()

	loginTest := dtos.LoginDTO{
		Email:    "test",
		Password: "test",
	}
	userTest := dtos.UserDTO{
		Name:  "test",
		Login: loginTest,
	}

	userOutput, err := userService.CreateUser(userTest)
	if err != nil {
		t.Fatal(err)
	}

	if userOutput == nil {
		t.Fatal("An error ocurred when create user")
	}
}

func TestCreateUserThatAlredyExists(t *testing.T) {
	userService := services.NewUserService()

	loginTest := dtos.LoginDTO{
		Email:    "CreateUserTest2",
		Password: "CreateUserTest2",
	}
	userTest := dtos.UserDTO{
		Name:  "CreateUserTest2",
		Login: loginTest,
	}

	_, err := userService.CreateUser(userTest)
	if err != nil {
		t.Fatal(err)
	}

	_, err = userService.CreateUser(userTest)
	if err == nil {
		t.Fatal("Has expected an error, because user alredy exist")
	}
}
