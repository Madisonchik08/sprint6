package service

import (
	"errors"
	"log"

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
		result := morse.ToText(data)
		log.Printf("Service: Converted Morse to Text: %q", result)
		return result, nil

	} else {
		result := morse.ToMorse(data)
		log.Printf("Service: Converted Text to Morse: %q", result)
		return result, nil
	}
}
