package integration

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/services"
	"testing"
)

func TestCreateAfiliate(t *testing.T) {
	userService := services.NewUserService()
	creatorService := services.NewCreatorSerivce()
	afiliateService := services.NewAfiliatedService()

	login := dtos.LoginDTO{
		Email:    "EmailAfiliated",
		Password: "PasswordAfiliated",
	}

	user := dtos.UserDTO{
		Name:  "UserAfiliate",
		Login: login,
	}

	creator := dtos.CreatorDTO{
		Name: "CreatorAfiliate",
	}

	afiliate := dtos.AfiliatedDTO{
		Name: "SimpleAfiliate",
	}

	userOuput, _ := userService.CreateUser(user)
	creatorService.CreateNewCreator(&creator, userOuput.UserId)
	validation := afiliateService.AddAfiliate(&afiliate, userOuput.UserId)

	if validation != nil {
		t.Fatal("An error ocurred, does not expired that afialite is not added")
	}
}

func TestCreateAfiliateThatAlredyExist(t *testing.T) {
	userService := services.NewUserService()
	creatorService := services.NewCreatorSerivce()
	afiliateService := services.NewAfiliatedService()

	login := dtos.LoginDTO{
		Email:    "EmailAfiliated2",
		Password: "PasswordAfiliated2",
	}

	user := dtos.UserDTO{
		Name:  "UserAfiliate2",
		Login: login,
	}

	creator := dtos.CreatorDTO{
		Name: "CreatorAfiliate2",
	}

	afiliate := dtos.AfiliatedDTO{
		Name: "SimpleAfiliate2",
	}

	userOuput, _ := userService.CreateUser(user)
	creatorService.CreateNewCreator(&creator, userOuput.UserId)
	validation := afiliateService.AddAfiliate(&afiliate, userOuput.UserId)

	if validation != nil {
		t.Fatal("An error ocurred, does not expired that afialite is not added")
	}

	validation = afiliateService.AddAfiliate(&afiliate, userOuput.UserId)
	if validation == nil {
		t.Fatal("An error not ocurred, expected that afiliate contains unique name")
	}
}

func TestCreateAfiliateThatUserIdIsInvalid(t *testing.T) {

	afiliateService := services.NewAfiliatedService()

	afiliate := dtos.AfiliatedDTO{
		Name: "SimpleAfiliate3",
	}

	validation := afiliateService.AddAfiliate(&afiliate, 0)

	if validation == nil {
		t.Fatal("An error not ocurred, expect the affiliate not to be created")
	}

}
