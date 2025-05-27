package numeralsystem

type ILexoNumeralSystem interface {
	GetBase() int
	GetPositiveChar() string
	GetNegativeChar() string
	GetRadixPointChar() string
	ToDigit(ch string) (int, error)
	ToChar(digit int) string
}
