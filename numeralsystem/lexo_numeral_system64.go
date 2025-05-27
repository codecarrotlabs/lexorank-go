package numeralsystem

import "strings"

type LexoNumeralSystem64 struct{}

func (LexoNumeralSystem64) GetBase() int {
	return 64
}

func (LexoNumeralSystem64) ToChar(val int) string {
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
	return string(chars[val])
}

func (LexoNumeralSystem64) ToValue(c rune) int {
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
	return strings.IndexRune(chars, c)
}