package unit

import (
	"hubla/desafiofullstack/models"
	"hubla/desafiofullstack/utils"
	"strconv"
	"testing"
)

func TestValidateTokenGeneratorJWT(t *testing.T) {
	authService := utils.NewAuth()

	input := &models.UserModel{
		UserId: 1,
	}

	token, err := authService.GenerateTokenJWT(input)
	if err != nil {
		t.Fatal(err.Error())
	}

	userIdReceived, ok := authService.ValidateToken(token)
	userId, _ := strconv.Atoi(userIdReceived)

	if userId != input.UserId && !ok {
		t.Fatal("Erro when validate token")
	}
}
