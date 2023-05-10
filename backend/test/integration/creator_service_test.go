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
	user := dtos.UserDTO{
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

func TestCreateCreatorWithInvalidUser(t *testing.T) {
	creatorService := services.NewCreatorSerivce()

	creator := dtos.CreatorDTO{
		Name: "Simple Name",
	}

	err := creatorService.CreateNewCreator(&creator, 0)
	if err == nil {
		t.Fatal("An error not ocurred, expected the creator not be created")
	}
}
