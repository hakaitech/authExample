package auth

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SECRET = "MyPasswordsSuperSecureSecret"

func GenerateToken() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SECRET))
	if err != nil {
		log.Println("Error in signing Token: ", err)
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(signedToken string) (claims *jwt.StandardClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwt.StandardClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, errors.New("claims invalid")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("Token expired, Pls Login again")
	}

	return claims, nil

}
