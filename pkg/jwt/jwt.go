package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hramov/tg-bot-admin/internal/config"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Claims struct {
	Id    string `json:"jti,omitempty"`
	Email string `json:"email"`
	Exp   int64  `json:"exp"`
}

func (c Claims) Valid() error { return nil } // TODO

const (
	AccessToken = iota
	RefreshToken
)

func CreateHashedPassword(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(plain string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		return false
	}
	return true
}

func CreateToken(id int, secret string) (string, string, error) {
	atClaims := Claims{}
	atClaims.Exp = time.Now().Add(config.JwtAccessTime).Unix()
	atClaims.Id = strconv.Itoa(id)

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}

	rtClaims := Claims{}
	rtClaims.Exp = time.Now().Add(config.JwtRefreshTime).Unix()
	rtClaims.Id = strconv.Itoa(id)

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := rt.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func VerifyToken(tokenString string, tokenType int) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	if claims["exp"] == nil {
		return nil, fmt.Errorf("token doesn't have exp value")
	}
	exp := int64(claims["exp"].(float64))
	if exp < time.Now().Unix() {
		return nil, fmt.Errorf("token is expired, use refresh token")
	}
	return token, nil
}

func TokenValid(tokenString string, tokenType int) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString, tokenType)
	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}

func CheckRefreshToken(t string) (int, error) {
	token, err := VerifyToken(t, 1)
	if err != nil {
		return 0, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	exp := int64(claims["exp"].(float64))
	if exp < time.Now().Unix() {
		return 0, fmt.Errorf("token is expired, use refresh token")
	}
	if claims["jti"] == nil {
		return 0, fmt.Errorf("token not valid")
	}
	rawId := claims["jti"].(string)
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func CreateClientId() string {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(1 * time.Nanosecond)
	const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ."
	b := make([]byte, 15)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func CreateClientSecret() (string, error) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(1 * time.Nanosecond)
	const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ."
	b := make([]byte, 15)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return CreateHashedPassword(string(b))
}

func CreateJWTSecret() (string, error) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(1 * time.Nanosecond)
	return "", nil
}

func CreateAuthCode() string {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(1 * time.Nanosecond)
	const letterBytes = "1234567890"
	b := make([]byte, 7)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
