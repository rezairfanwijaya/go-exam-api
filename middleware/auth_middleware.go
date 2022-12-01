package middleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rezairfanwijaya/go-exam-api.git/helper"
	"github.com/rezairfanwijaya/go-exam-api.git/user"
)

type IAuthService interface {
	GenerateJWT(user user.User) (string, error)
	ValidasiJWT(jwtToken string) (*jwt.Token, error)
}

type tokenJWT struct{}

func NewServiceAuth() *tokenJWT {
	return &tokenJWT{}
}

func (t *tokenJWT) GenerateJWT(user user.User) (string, error) {
	// set payload jwt
	claims := jwt.MapClaims{}
	claims["userID"] = user.ID
	claims["email"] = user.Email
	claims["expired"] = time.Now().Add(5 * time.Minute).Unix()

	// apply algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// ambil key secret
	env, err := helper.GetEnv(".env")
	if err != nil {
		return "", err
	}

	// tanda tangan
	jwt, err := token.SignedString([]byte(env["KEY"]))
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (t *tokenJWT) ValidasiJWT(jwtToken string) (*jwt.Token, error) {
	// parsing jwt
	myToken, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// cek algorithm
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		// get secret key
		env, err := helper.GetEnv(".env")
		if err != nil {
			return nil, err
		}

		return []byte(env["KEY"]), nil
	})

	if err != nil {
		return nil, err
	}

	return myToken, nil
}
