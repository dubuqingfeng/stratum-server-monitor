package dcr

import "strconv"

func ParseHeight(coinbaseTx1 string) (height int64, err error) {
	heightHex := coinbaseTx1[190:192] + coinbaseTx1[188:190] + coinbaseTx1[186:188] + coinbaseTx1[184:186]
	height, err = strconv.ParseInt(heightHex, 16, 32)
	return
}
