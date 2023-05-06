package unit

import (
	"hubla/desafiofullstack/dtos"
	"hubla/desafiofullstack/utils"
	"testing"
)

func TestValidateTokenGeneratorJWT(t *testing.T) {
	authService := utils.NewAuth()

	input := &dtos.LoginDTO{
		Email:    "simpleExampleTest@hotmail.com",
		Password: "simplePassword123",
	}

	token, err := authService.GenerateTokenJWT(input)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedEmail, ok := authService.ValidateToken(token)

	if expectedEmail != input.Email && !ok {
		t.Fatal("Erro when validate token")
	}
}
