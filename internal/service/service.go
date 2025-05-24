package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Conversion(value string) (string, error) {
	alpha := "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЭЮЯ"

	if value == "" {
		return "", fmt.Errorf("string is empty")
	}

	if strings.ContainsAny(value, alpha) || strings.ContainsAny(value, strings.ToLower(alpha)) {
		return morse.ToMorse(value), nil
	} else {
		return morse.ToText(value), nil
	}
}
