package integration

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/services"
	"testing"
)

func TestCreateCreator(t *testing.T) {
	userService := services.NewUserService()
	creatorService := services.NewCreatorSerivce()

	login := dtos.LoginDTO{
		Email:    "test_3",
		Password: "test_3",
	}
	user := dtos.CreateUseDTO{
		Name:  "test",
		Login: login,
	}

	creator := dtos.CreatorDTO{
		Name: "Simple Name",
	}

	usertOutput, err := userService.CreateUser(user)
	if err != nil {
		t.Fatal("Doesn't expire erro when create user")
	}

	err = creatorService.CreateNewCreator(&creator, usertOutput.UserId)
	if err != nil {
		t.Fatal("Doesn't expire erro when create an creator")
	}
}
