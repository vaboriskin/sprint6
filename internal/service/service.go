package service

import (
	"errors"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertMorseCode(input string) (string, error) {
	isMorse := true

	for _, ch := range input {
		if ch == ' ' || ch == '\t' || ch == '\n' {
			continue
		}
		if ch != '.' && ch != '-' {
			isMorse = false
			break
		}
	}

	if isMorse {

		return morse.ToText(input), nil
	}

	return morse.ToMorse(input), nil
}

var (
	ErrInvalidInput = errors.New("invalid input: not text or morse code")
)
