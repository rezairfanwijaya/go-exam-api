package helper

import (
	"errors"

	"github.com/joho/godotenv"
)

// ? function untuk mengambil value env
func GetEnv(path string) (map[string]string, error) {
	env, err := godotenv.Read(path)
	if err != nil {
		return nil, errors.New("missing path, file doesnt exist")
	}
	return env, nil
}
