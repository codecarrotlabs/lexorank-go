package numeralsystem

import (
    "fmt"
)

type LexoNumeralSystem10 struct{}

func NewLexoNumeralSystem10() *LexoNumeralSystem10 {
    return &LexoNumeralSystem10{}
}

func (l *LexoNumeralSystem10) GetBase() int {
    return 10
}

func (l *LexoNumeralSystem10) GetPositiveChar() rune {
    return '+'
}

func (l *LexoNumeralSystem10) GetNegativeChar() rune {
    return '-'
}

func (l *LexoNumeralSystem10) GetRadixPointChar() rune {
    return '.'
}

func (l *LexoNumeralSystem10) ToDigit(ch rune) (int, error) {
    if ch >= '0' && ch <= '9' {
        return int(ch - '0'), nil
    }
    return 0, fmt.Errorf("not valid digit: %c", ch)
}

func (l *LexoNumeralSystem10) ToChar(digit int) (rune, error) {
    if digit < 0 || digit > 9 {
        return 0, fmt.Errorf("digit out of range: %d", digit)
    }
    return rune(digit + '0'), nil
}