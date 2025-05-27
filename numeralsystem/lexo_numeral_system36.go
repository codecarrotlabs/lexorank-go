package numeralsystem

import "strings"

type LexoNumeralSystem36 struct{}

func (LexoNumeralSystem36) GetBase() int {
	return 36
}

func (LexoNumeralSystem36) ToChar(val int) string {
	const chars = "0123456789abcdefghijklmnopqrstuvwxyz"
	return string(chars[val])
}

func (LexoNumeralSystem36) ToValue(c rune) int {
	const chars = "0123456789abcdefghijklmnopqrstuvwxyz"
	return strings.IndexRune(chars, c)
}