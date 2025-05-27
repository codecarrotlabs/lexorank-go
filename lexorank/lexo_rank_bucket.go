package lexorank

import "fmt"

type LexoRankBucket string

const (
	Bucket0 LexoRankBucket = "0"
	Bucket1 LexoRankBucket = "1"
	Bucket2 LexoRankBucket = "2"
)

func ParseBucket(s string) (LexoRankBucket, error) {
	switch s {
	case "0":
		return Bucket0, nil
	case "1":
		return Bucket1, nil
	case "2":
		return Bucket2, nil
	default:
		return "", fmt.Errorf("invalid bucket: %s", s)
	}
}

func (b LexoRankBucket) String() string {
	return string(b)
}