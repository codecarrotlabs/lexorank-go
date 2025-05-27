package lexorank

import (
	"errors"
	"fmt"
	"lexorank-go/lexorank"
	"lexorank-go/numeralsystem"
	"strings"
)

type LexoRank struct {
	Bucket LexoRankBucket
	Decimal *LexoDecimal
}

func NewLexoRank(str string) (*LexoRank, error) {
	parts := strings.Split(str, "|")
	if len(parts) != 2 {
		return nil, errors.New("invalid LexoRank string")
	}
	bucket, err := ParseBucket(parts[0])
	if err != nil {
		return nil, err
	}
	decimal, err := ParseDecimal(parts[1])
	if err != nil {
		return nil, err
	}
	return &LexoRank{
		Bucket:  bucket,
		Decimal: decimal,
	}, nil
}

func (l *LexoRank) Next() (*LexoRank, error) {
	nextDecimal := l.Decimal.Next()
	return &LexoRank{
		Bucket:  l.Bucket,
		Decimal: nextDecimal,
	}, nil
}

func (l *LexoRank) String() string {
	return fmt.Sprintf("%s|%s", l.Bucket.String(), l.Decimal.String())
}