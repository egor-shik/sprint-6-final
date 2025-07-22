package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/egor-shik/sprint-6-final/pkg/morse"
)

// ConvertResult содержит результат конвертации и информацию о типе исходных данных
type ConvertResult struct {
	Result   string
	WasMorse bool
	Original string
}

// AutoConvert автоматически определяет тип ввода (текст или код Морзе)
// и выполняет соответствующую конвертацию
func AutoConvert(input string) (ConvertResult, error) {
	if strings.TrimSpace(input) == "" {
		return ConvertResult{}, errors.New("пустой ввод")
	}

	if isMorseCode(input) {
		text, err := morse.ToText(input)
		if err != nil {
			return ConvertResult{}, err
		}
		return ConvertResult{
			Result:   text,
			WasMorse: true,
			Original: input,
		}, nil
	}

	morseCode, err := morse.ToMorse(input)
	if err != nil {
		return ConvertResult{}, err
	}
	return ConvertResult{
		Result:   morseCode,
		WasMorse: false,
		Original: input,
	}, nil
}

// isMorseCode проверяет, является ли строка кодом Морзе
func isMorseCode(s string) bool {
	// Допустимые символы в коде Морзе: точка, тире, пробел и слэш
	morseChars := ".-/ "

	// Сначала проверяем, что нет "обычных" букв/цифр
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return false
		}
	}

	// Затем проверяем, что все символы - допустимые символы Морзе
	for _, r := range s {
		if !strings.ContainsRune(morseChars, r) {
			return false
		}
	}

	// Минимальная проверка структуры (хотя бы одна точка или тире)
	return strings.ContainsAny(s, ".-")
}
