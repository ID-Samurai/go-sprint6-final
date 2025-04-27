package service

import (
	"errors"

	"github.com/ID-Samurai/go-sprint6-final/pkg/morse"
)

func isMorse(s string) bool {
	for _, r := range s {
		if r != '.' && r != ' ' && r != '-' {
			return false
		}
	}
	return true
}
func AutoConversion(s string) (str string, err error) {
	if len(s) == 0 {
		return "", errors.New("Error: an empty string cannot be converted")
	}
	if isMorse(s) == true {
		return morse.ToText(s), nil
	} else {
		return morse.ToMorse(s), nil
	}
}
