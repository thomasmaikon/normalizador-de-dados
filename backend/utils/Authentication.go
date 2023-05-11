package utils

import (
	"fmt"
	"hubla/desafiofullstack/models"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type auth struct {
}

func NewAuth() *auth {
	return &auth{}
}

func (auth *auth) GenerateTokenJWT(user *models.UserModel) (string, error) {
	key := []byte(os.Getenv("secretkey"))

	id := strconv.Itoa(user.UserId)

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   id,
		})

	result, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (auth *auth) ValidateToken(token string) (string, bool) {

	result, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("secretkey")), nil
	})

	if err != nil {
		return "", false
	}

	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		userID := claims["sub"].(string)
		return userID, true
	} else {
		return "", false
	}

}

func IsAuthorized(ctx *gin.Context) {
	BEARER_SCHEMA := "Bearer"

	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA)+1:]

	userID, isValid := NewAuth().ValidateToken(tokenString)

	if !isValid {
		ctx.AbortWithStatus(401)
		return
	}
	ctx.Set("userID", userID)
	ctx.Next()
}
