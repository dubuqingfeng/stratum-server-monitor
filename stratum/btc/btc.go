package btc

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func ParseHeight(coinbaseTx1 string) (height int64, err error) {
	if len(coinbaseTx1) < 92 {
		return 0, errors.New("error coinbase.")
	}
	heightHex := coinbaseTx1[90:92] + coinbaseTx1[88:90] + coinbaseTx1[86:88]
	height, err = strconv.ParseInt(heightHex, 16, 32)
	return
}

func ParsePrevHash(hash string) string {
	var splits []string
	for i := 0; i < len(hash); i += 8 {
		var appendStr string
		if (i + 8) > len(hash) {
			appendStr = hash[i:]
		} else {
			appendStr = hash[i : i+8]
		}
		splits = append(splits, appendStr)
	}
	// reverse slice
	for from, to := 0, len(splits)-1; from < to; from, to = from+1, to-1 {
		splits[from], splits[to] = splits[to], splits[from]
	}
	return strings.Join(splits, "")
}
