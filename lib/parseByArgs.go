package lib

import (
	"errors"
)

// function untuk mencari data dari argumen
func ParseByArgs(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("Silahkan masukan argumen")
	}
	return args[1], nil
}