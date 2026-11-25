package service

import(
	"strings"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"fmt"
)

func CheckString(text string) (string, error) {
	var result string
	if len(text) == 0 || strings.TrimSpace(text) == ""{
		err := fmt.Errorf("empty string")
		return "", err
	}
	if strings.IndexFunc(text, func(r rune) bool {return r != '.' && r != '-' && r != ' '}) == -1 {
		result = morse.ToText(text)
	} else {
		result = morse.ToMorse(text)
	}
	return result, nil
}
