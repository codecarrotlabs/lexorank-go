package lexorank

import (
	"math/big"
	"strings"
)

// LexoInteger is a wrapper for big.Int with scale and base.
type LexoInteger struct {
	Value *big.Int
	Scale int
	Base  int
}

func NewLexoIntegerFromString(s string, base int) (*LexoInteger, error) {
	val, ok := new(big.Int).SetString(s, base)
	if !ok {
		return nil, ErrInvalidInteger
	}
	return &LexoInteger{
		Value: val,
		Scale: len(s),
		Base:  base,
	}, nil
}

func NewLexoInteger(val *big.Int, scale, base int) *LexoInteger {
	return &LexoInteger{
		Value: new(big.Int).Set(val),
		Scale: scale,
		Base:  base,
	}
}

func (li *LexoInteger) String() string {
	s := strings.ToLower(li.Value.Text(li.Base))
	for len(s) < li.Scale {
		s = "0" + s
	}
	return s
}

var ErrInvalidInteger = &ParseError{"invalid integer string"}