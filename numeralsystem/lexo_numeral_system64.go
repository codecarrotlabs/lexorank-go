package numeralsystem

import (
    "fmt"
)

type LexoNumeralSystem64 struct {
    digits []rune
}

func NewLexoNumeralSystem64() *LexoNumeralSystem64 {
    return &LexoNumeralSystem64{
        digits: []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ^_abcdefghijklmnopqrstuvwxyz"),
    }
}

func (l *LexoNumeralSystem64) GetBase() int {
    return 64
}

func (l *LexoNumeralSystem64) GetPositiveChar() rune {
    return '+'
}

func (l *LexoNumeralSystem64) GetNegativeChar() rune {
    return '-'
}

func (l *LexoNumeralSystem64) GetRadixPointChar() rune {
    return ':'
}

func (l *LexoNumeralSystem64) ToDigit(ch rune) (int, error) {
    switch {
    case ch >= '0' && ch <= '9':
        return int(ch - '0'), nil
    case ch >= 'A' && ch <= 'Z':
        return int(ch-'A') + 10, nil
    case ch == '^':
        return 36, nil
    case ch == '_':
        return 37, nil
    case ch >= 'a' && ch <= 'z':
        return int(ch-'a') + 38, nil
    default:
        return 0, fmt.Errorf("not valid digit: %c", ch)
    }
}

func (l *LexoNumeralSystem64) ToChar(digit int) (rune, error) {
    if digit < 0 || digit >= len(l.digits) {
        return 0, fmt.Errorf("digit out of range: %d", digit)
    }
    return l.digits[digit], nil
}