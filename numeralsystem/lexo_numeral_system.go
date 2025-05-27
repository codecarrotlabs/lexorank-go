package numeralsystem

type LexoNumeralSystem interface {
	GetBase() int
	ToChar(val int) string
	ToValue(c rune) int
}