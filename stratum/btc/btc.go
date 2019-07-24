package btc

import (
	"github.com/pkg/errors"
	"strconv"
)

func ParseHeight(coinbaseTx1 string) (height int64, err error) {
	if len(coinbaseTx1) < 92 {
		return 0, errors.New("error coinbase.")
	}
	heightHex := coinbaseTx1[90:92] + coinbaseTx1[88:90] + coinbaseTx1[86:88]
	height, err = strconv.ParseInt(heightHex, 16, 32)
	return
}
