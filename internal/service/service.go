package service

import (
	"errors"
	"sprint-6-final/pkg/morse"
	"strings"
	"unicode"
)

// ConvertResult содержит результат конвертции
type ConvertResult struct {
	Result   string // Конвертированный результат
	WasMorse bool   // Флаг, указывюащий был ли ввод кодом морзе
	Original string // Изначальный ыыод
}

func AutoConvert(input string) (ConvertResult, error) {
	if strings.TrimSpace(input) == "" {
		return ConvertResult{}, errors.New("пустой ввод")
	}

	if isMorseCode(input) {
		text := morse.ToText(input)
		return ConvertResult{
			Result:   text,
			WasMorse: true,
			Original: input,
		}, nil
	}

	morseCode := morse.ToMorse(input)
	return ConvertResult{
		Result:   morseCode,
		WasMorse: false,
		Original: input,
	}, nil
}

// isMorseCode проверяет, является ли строка кодом морзе
func isMorseCode(s string) bool {
	morseChars := ".-/ "

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return false
		}
	}

	for _, r := range s {
		if !strings.ContainsRune(morseChars, r) {
			return false
		}
	}

	return strings.ContainsAny(s, ".-")
}
