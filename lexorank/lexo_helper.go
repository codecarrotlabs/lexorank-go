package lexorank

import "math/big"

// LexoHelper provides utility functions for LexoRank.
type LexoHelper struct{}

func (LexoHelper) Max(a, b *big.Int) *big.Int {
	if a.Cmp(b) >= 0 {
		return a
	}
	return b
}

func (LexoHelper) Min(a, b *big.Int) *big.Int {
	if a.Cmp(b) <= 0 {
		return a
	}
	return b
}