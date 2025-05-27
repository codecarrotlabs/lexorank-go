package numeralsystem

import "strings"

type LexoNumeralSystem10 struct{}

func (LexoNumeralSystem10) GetBase() int {
	return 10
}

func (LexoNumeralSystem10) ToChar(val int) string {
	const chars = "0123456789"
	return string(chars[val])
}

func (LexoNumeralSystem10) ToValue(c rune) int {
	const chars = "0123456789"
	return strings.IndexRune(chars, c)
}