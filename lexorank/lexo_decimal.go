package lexorank

import (
	"lexorank-go/numeralsystem"
	"math/big"
	"strings"
)

type LexoDecimal struct {
	Value  *big.Int
	System numeralsystem.LexoNumeralSystem
	Scale  int
}

func ParseDecimal(s string) (*LexoDecimal, error) {
	sys := numeralsystem.LexoNumeralSystem36{}
	val, ok := new(big.Int).SetString(s, 36)
	if !ok {
		return nil, ErrInvalidDecimal
	}
	return &LexoDecimal{
		Value:  val,
		System: sys,
		Scale:  len(s),
	}, nil
}

func (d *LexoDecimal) Next() *LexoDecimal {
	next := new(big.Int).Add(d.Value, big.NewInt(1))
	return &LexoDecimal{
		Value:  next,
		System: d.System,
		Scale:  d.Scale,
	}
}

func (d *LexoDecimal) String() string {
	str := strings.ToLower(d.Value.Text(36))
	for len(str) < d.Scale {
		str = "0" + str
	}
	return str
}

var ErrInvalidDecimal = &ParseError{"invalid decimal string"}