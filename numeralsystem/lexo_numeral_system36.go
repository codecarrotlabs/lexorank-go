package numeralsystem

import (
    "fmt"
)

type LexoNumeralSystem36 struct {
    digits []rune
}

func NewLexoNumeralSystem36() *LexoNumeralSystem36 {
    return &LexoNumeralSystem36{
        digits: []rune("0123456789abcdefghijklmnopqrstuvwxyz"),
    }
}

func (l *LexoNumeralSystem36) GetBase() int {
    return 36
}

func (l *LexoNumeralSystem36) GetPositiveChar() rune {
    return '+'
}

func (l *LexoNumeralSystem36) GetNegativeChar() rune {
    return '-'
}

func (l *LexoNumeralSystem36) GetRadixPointChar() rune {
    return ':'
}

func (l *LexoNumeralSystem36) ToDigit(ch rune) (int, error) {
    switch {
    case ch >= '0' && ch <= '9':
        return int(ch - '0'), nil
    case ch >= 'a' && ch <= 'z':
        return int(ch-'a') + 10, nil
    default:
        return 0, fmt.Errorf("not valid digit: %c", ch)
    }
}

func (l *LexoNumeralSystem36) ToChar(digit int) (rune, error) {
    if digit < 0 || digit >= len(l.digits) {
        return 0, fmt.Errorf("digit out of range: %d", digit)
    }
    return l.digits[digit], nil
}