package ckb

import (
	"github.com/pkg/errors"
	"math"
)

func ParseHeight(height float64) (int64, error) {
	if height >= math.MaxInt64 || height <= math.MinInt64 {
		return 0, errors.New("More than an int64 range")
	}
	return int64(height), nil
}

func ParsePrevHash(hash string) string {
	return hash
}
