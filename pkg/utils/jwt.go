package utils

import (
	"advent-calendar/internal/config"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
}

func NewJWT(id uint, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Config.SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func NewRefreshToken() (string, error) {
	b := make([]byte, 32)
	r := rand.New(rand.NewSource(time.Now().Unix()))

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func VerifyToken(token string) (Claims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.SECRET), nil
	})
	if err != nil {
		return Claims{}, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return Claims{}, errors.New("Ошибка при парсинге токена")
	}

	return Claims{ID: uint(claims["id"].(float64)), Role: claims["role"].(string)}, nil
}

func GetBearerToken(c *fiber.Ctx, field string) (string, error) {
	authHeader := c.Get(field)

	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		return strings.Split(authHeader, " ")[1], nil
	}

	return "", errors.New("Отсутствует токен авторизации")

}
