package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert функиця принимает текст и возвращает конвертацию
func Convert(data string) (string, error) {
	input := strings.TrimSpace(data)

	if input == "" {
		return "", morse.ErrNoEncoding{}
	}

	// Определяем тип ввода по наличию морзе-символов
	if isMorseCode(input) {
		return morseToText(input)
	}
	return textToMorse(input)
}

// isMorseCode функция проверяет является ли входной текст морзой
func isMorseCode(s string) bool {
	return !strings.ContainsFunc(s, func(r rune) bool {
		return !isValidMorseChar(r)
	})
}

// isValidMorseChar функция проверяет содержит ли текст символы морзы
func isValidMorseChar(r rune) bool {
	switch r {
	case '.', '-', ' ', '/', '(', ')', '"', '=', '+', '@', '_':
		return true
	}
	return false
}

// textToMorse функция конвертирует текст в морзе
func textToMorse(text string) (string, error) {
	converted := morse.ToMorse(text)

	if strings.Contains(converted, string([]rune{0xFFFD})) {
		return "", fmt.Errorf("text has unknown symbols")
	}

	return converted, nil
}

// morseToText функция конвертирует морзе в текст
func morseToText(mrsStr string) (string, error) {
	converted := morse.ToText(mrsStr)

	if strings.Contains(converted, string([]rune{0xFFFD})) {
		return "", fmt.Errorf("morse has unknown symbols")
	}

	return converted, nil
}
