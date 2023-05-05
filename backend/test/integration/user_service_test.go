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
	userTest := dtos.CreateUseDTO{
		Name:  "test",
		Login: loginTest,
	}

	err := userService.CreateUser(userTest)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateUserThatAlredyExists(t *testing.T) {
	userService := services.NewUserService()

	loginTest := dtos.LoginDTO{
		Email:    "test2",
		Password: "test",
	}
	userTest := dtos.CreateUseDTO{
		Name:  "test2",
		Login: loginTest,
	}

	err := userService.CreateUser(userTest)
	if err != nil {
		t.Fatal(err)
	}

	err = userService.CreateUser(userTest)
	if err == nil {
		t.Fatal("Has expected an error, because user alredy exist")
	}
}
