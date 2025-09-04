package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ProcessData(data string) (string, error) {
	if data == "" {
		return "", errors.New("data is empty")
	}
	isMorse := true
	for _, r := range data {
		if r != '.' && r != '-' && r != '/' && r != ' ' {
			isMorse = false
			break
		}
	}
	if isMorse {
		return morse.ToText(data), nil

	} else {
		normalized := morse.ToText(data)
		return morse.ToText(normalized), nil
	}
}
func normalizeText(text string) string {
	var result []rune
	for _, r := range strings.ToUpper(text) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			result = append(result, r)
		}
	}
	return string(result)
}
