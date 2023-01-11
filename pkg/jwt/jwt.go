package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

type Claims struct {
	Id          string      `json:"jti,omitempty"`
	Email       string      `json:"email"`
	Exp         int64       `json:"exp"`
	Permissions Permissions `json:"permissions"`
}

type Permissions struct {
	Admin bool     `json:"admin,omitempty"`
	Scope []string `json:"scope,omitempty"`
}

func (c Claims) Valid() error { return nil } // TODO

func CreateToken(id int, perm Permissions, accessSecret, refreshSecret string, accessTtl, refreshTtl time.Duration) (string, string, error) {
	atClaims := Claims{}
	atClaims.Exp = time.Now().Add(accessTtl).Unix()
	atClaims.Id = strconv.Itoa(id)
	atClaims.Permissions = perm

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString([]byte(accessSecret))
	if err != nil {
		return "", "", err
	}

	rtClaims := Claims{}
	rtClaims.Exp = time.Now().Add(refreshTtl).Unix()
	rtClaims.Id = strconv.Itoa(id)

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := rt.SignedString([]byte(refreshSecret))
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func GetClaims(tokenString string, secret string) (jwt.MapClaims, error) {
	token, err := verifyToken(tokenString, secret)
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}

func CheckRefreshToken(t string, secret string) (int, error) {
	token, err := verifyToken(t, secret)
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

func verifyToken(tokenString string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
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
